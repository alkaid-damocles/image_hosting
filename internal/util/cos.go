package util

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
)

const (
	CosSecretID  = ""
	CosSecretKey = ""
	CosBucketURL = ""
	SystemType   = "linux"
)

func UploadToCos(filePath string) string {
	ctx := context.Background()
	u, _ := url.Parse(CosBucketURL)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: CosSecretID, // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: CosSecretKey, // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参见 https://cloud.tencent.com/document/product/598/37140
		},
	})

	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentDisposition: "attachment",
		},
	}

	t := time.Now()
	// 使用 PUtFromFile 上传本地文件到 COS
	fileName := FileNameFormat(filePath)
	key := fmt.Sprintf("%d/%d/%d/%s", t.Year(), t.Month(), t.Day(), fileName)
	_, err := client.Object.PutFromFile(ctx, key, filePath, opt)
	if err != nil {
		panic(err)
	}

	getOpt := &cos.ObjectGetOptions{
		ResponseContentDisposition: fmt.Sprintf("attachment; filename=%s", fileName),
	}
	presignedURL, err := client.Object.GetPresignedURL(ctx, http.MethodGet, key, CosSecretID, CosSecretKey, 20*365*24*(time.Hour), getOpt)
	if err != nil {
		log.Fatalln(err)
		return ""
	}
	return fmt.Sprint(presignedURL)
}

func FileNameFormat(filePath string) string {
	fileName := filePath
	switch SystemType {
	case "linux":
		s := strings.Split(filePath, "/")
		if len(s) > 1 {
			fileName = s[len(s)-1]
		}
	case "windows":
		s := strings.Split(filePath, "\\")
		if len(s) > 1 {
			fileName = s[len(s)-1]
		}
	}
	return fileName
}
