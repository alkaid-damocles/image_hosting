package main

import (
	"fmt"
	"os"

	"github.com/alkaid-damocles/image_hosting/internal/util"
)

func main() {
	pictrueURLs := os.Args
	for index, pictrueURL := range pictrueURLs {
		if (index) == 0 {
			continue
		}
		url := util.UploadToCos(pictrueURL)
		fmt.Println(url)
	}
}
