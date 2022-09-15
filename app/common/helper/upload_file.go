package helper

import (
	"bytes"
	"context"
	"core/app/common/vars"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var client *cos.Client

// 获取 *cos.Client
func getClient(server *vars.CosConfig) *cos.Client {
	if client != nil {
		return client
	}
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(server.Server)
	b := &cos.BaseURL{BucketURL: u}
	client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  server.SecretID,
			SecretKey: server.SecretKey,
		},
	})
	return client
}

// 根据http请求上传文件，返回路径
func UploadFile(CosConfigObject *vars.CosConfig, r *http.Request) (string, error) {

	file, fileHeader, err := r.FormFile("file")

	if err != nil {
		return "", err
	}

	key := GenerateUuid() + path.Ext(fileHeader.Filename)
	client := getClient(CosConfigObject)
	_, err = client.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		return "", err
	}
	return CosConfigObject.Server + "/" + key, nil
}

// 分片上传初始化
func ChunkInit(CosConfigObject *vars.CosConfig, ext string) (key, uploadId string, err error) {
	client := getClient(CosConfigObject)
	key = GenerateUuid() + ext
	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), key, nil)
	if err != nil {
		return
	}
	uploadId = v.UploadID
	return
}

//分片上传
func ChunkUpload(CosConfigObject *vars.CosConfig, r *http.Request) (eTag string, err error) {
	client := getClient(CosConfigObject)
	file, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	key := r.PostForm.Get("key")
	upload_id := r.PostForm.Get("upload_id")
	part_number, _ := strconv.Atoi(r.PostForm.Get("part_number"))

	buf := bytes.NewBuffer(nil)
	io.Copy(buf, file)

	resp, err := client.Object.UploadPart(
		context.Background(), key, upload_id, part_number, bytes.NewReader(buf.Bytes()), nil,
	)
	if err != nil {
		return "", err
	}
	eTag = resp.Header.Get("ETag")
	eTag = strings.Trim(eTag, "\"")
	return
}

// 分片上传完成
// func ChunkSuccess(req *types.FileChunkSuccessRequest) (err error) {
// 	client := getClient()
// 	opt := &cos.CompleteMultipartUploadOptions{}
// 	for _, v := range req.Parts {
// 		opt.Parts = append(opt.Parts, cos.Object{
// 			PartNumber: v.PartNumber, ETag: v.ETag},
// 		)
// 	}

// 	_, _, err = client.Object.CompleteMultipartUpload(
// 		context.Background(), req.Key, req.UploadId, opt,
// 	)
// 	return
// }
