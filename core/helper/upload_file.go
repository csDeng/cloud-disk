package helper

import (
	"context"
	"net/http"
	"net/url"
	"path"

	"github.com/tencentyun/cos-go-sdk-v5"
)

// 根据http请求上传文件，返回路径
func UploadFile(r *http.Request) (string, error) {
	server := CosConfigObject
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(server.Server)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  server.SecretID,
			SecretKey: server.SecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		return "", err
	}

	key := GenerateUuid() + path.Ext(fileHeader.Filename)

	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return "", err
	}
	return CosConfigObject.Server + "/" + key, nil
}
