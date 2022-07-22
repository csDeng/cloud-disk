package test

import (
	"bytes"
	"context"
	"core/core/helper"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func TestUploadFileByPath(t *testing.T) {
	server := helper.CosConfigObject
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(server.Server)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: server.SecretID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: server.SecretKey,
		},
	})

	key := "exampleobject.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/0.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestUploadByReader(t *testing.T) {
	server := helper.CosConfigObject
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(server.Server)
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: server.SecretID,
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: server.SecretKey,
		},
	})

	// Case1 使用 Put 上传对象
	key := "exampleobject2.jpg"
	f, err := os.ReadFile("./img/0.jpg")
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
			XCosACL: "private",
		},
	}
	_, err = client.Object.Put(context.Background(), key, bytes.NewReader(f), opt)
	if err != nil {
		panic(err)
	}

	// Case 2 使用 PUtFromFile 上传本地文件到COS
	// filepath := "./test"
	// _, err = client.Object.PutFromFile(context.Background(), key, filepath, opt)
	// if err != nil {
	// 	panic(err)
	// }

	// // Case 3 上传 0 字节文件, 设置输入流长度为 0
	// _, err = client.Object.Put(context.Background(), key, strings.NewReader(""), nil)
	// if err != nil {
	// 	// ERROR
	// }
}
