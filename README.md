# image_hosting
使用腾讯云cos 和 golang实现的图床

填写自己的腾讯云cos  id  key 存储桶配置, 然后编译就可


让 go build 生成的可执行文件对 Mac、linux、Windows 平台一致
让 go build 生成的可执行文件对 Mac、linux、Windows 平台一致
 

要做到这一点，使用的是交叉编译选项。
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

确定目标机器的系统和架构，在运行 go build 的环境中，运行命令之前指定好相应的参数值。
