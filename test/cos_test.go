package test

import (
	"bytes"
	"context"
	"core/core/helper"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func getCosClient() *cos.Client {
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
	return client

}

func TestUploadFileByPath(t *testing.T) {
	client := getCosClient()

	key := "exampleobject.jpg"

	_, _, err := client.Object.Upload(
		context.Background(), key, "./img/0.jpg", nil,
	)
	if err != nil {
		panic(err)
	}
}

func TestUploadByReader(t *testing.T) {
	client := getCosClient()

	// Case1 使用 Put 上传对象
	key := "exampleobject2.jpg"
	f, err := os.ReadFile("./img/0.jpg")
	if err != nil {
		t.Fatal(err)
	}
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

var UploadId = "165928464019e8a6c471b83b57772f02c544d38b049de0d361ac4b83687a867b39a48a3616"
var Key = "test123"
var ETag string

// 初始化分片上传
func TestChunkInit(t *testing.T) {
	client := getCosClient()

	// 可选opt,如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
	v, _, err := client.Object.InitiateMultipartUpload(context.Background(), Key, nil)
	if err != nil {
		panic(err)
	}
	UploadId = v.UploadID
	fmt.Println(UploadId) // 1658943071d1960bf5540ab0be53ed65ebea8e319d8ee7275511d4ba3629b0bd8393e11b6c
}

// 文件分片上传COS
func TestChunkUpload(t *testing.T) {
	client := getCosClient()
	f, err := os.ReadFile("84_chunk")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("id=", UploadId)
	// opt可选
	resp, err := client.Object.UploadPart(
		context.Background(), Key, UploadId, 1, bytes.NewReader(f), nil,
	)
	if err != nil {
		panic(err)
	}
	PartETag := resp.Header.Get("ETag") // f73ac6d59f0f1228dd44d3139cec8e00
	fmt.Println(PartETag)

}

// 查询已上传分片
func TestCheckUpload(t *testing.T) {
	client := getCosClient()
	res := &cos.ListMultipartUploadsResult{
		XMLName: xml.Name{
			Space: "",
			Local: "",
		},
		Bucket:             "",
		EncodingType:       "",
		KeyMarker:          "",
		UploadIDMarker:     "",
		NextKeyMarker:      "",
		NextUploadIDMarker: "",
		MaxUploads:         0,
		IsTruncated:        false,
		Uploads: []struct {
			Key          string
			UploadID     string "xml:\"UploadId\""
			StorageClass string
			Initiator    *cos.Initiator
			Owner        *cos.Owner
			Initiated    string
		}{},
		Prefix:         "",
		Delimiter:      "",
		CommonPrefixes: []string{},
	}
	resp := &cos.Response{
		Response: &http.Response{
			Status:     "",
			StatusCode: 0,
			Proto:      "",
			ProtoMajor: 0,
			ProtoMinor: 0,
			Header: map[string][]string{
				"": {},
			},
			Body:             nil,
			ContentLength:    0,
			TransferEncoding: []string{},
			Close:            false,
			Uncompressed:     false,
			Trailer: map[string][]string{
				"": {},
			},
			Request: &http.Request{
				Method: "",
				URL: &url.URL{
					Scheme:      "",
					Opaque:      "",
					User:        &url.Userinfo{},
					Host:        "",
					Path:        "",
					RawPath:     "",
					ForceQuery:  false,
					RawQuery:    "",
					Fragment:    "",
					RawFragment: "",
				},
				Proto:      "",
				ProtoMajor: 0,
				ProtoMinor: 0,
				Header: map[string][]string{
					"": {},
				},
				Body:             nil,
				GetBody:          func() (io.ReadCloser, error) { panic("not implemented") },
				ContentLength:    0,
				TransferEncoding: []string{},
				Close:            false,
				Host:             "",
				Form: map[string][]string{
					"": {},
				},
				PostForm: map[string][]string{
					"": {},
				},
				MultipartForm: &multipart.Form{
					Value: map[string][]string{
						"": {},
					},
					File: map[string][]*multipart.FileHeader{
						"": {},
					},
				},
				Trailer: map[string][]string{
					"": {},
				},
				RemoteAddr: "",
				RequestURI: "",
				TLS: &tls.ConnectionState{
					Version:                     0,
					HandshakeComplete:           false,
					DidResume:                   false,
					CipherSuite:                 0,
					NegotiatedProtocol:          "",
					NegotiatedProtocolIsMutual:  false,
					ServerName:                  "",
					PeerCertificates:            []*x509.Certificate{},
					VerifiedChains:              [][]*x509.Certificate{},
					SignedCertificateTimestamps: [][]byte{},
					OCSPResponse:                []byte{},
					TLSUnique:                   []byte{},
				},
				Cancel:   make(<-chan struct{}),
				Response: &http.Response{},
			},
			TLS: &tls.ConnectionState{
				Version:                     0,
				HandshakeComplete:           false,
				DidResume:                   false,
				CipherSuite:                 0,
				NegotiatedProtocol:          "",
				NegotiatedProtocolIsMutual:  false,
				ServerName:                  "",
				PeerCertificates:            []*x509.Certificate{},
				VerifiedChains:              [][]*x509.Certificate{},
				SignedCertificateTimestamps: [][]byte{},
				OCSPResponse:                []byte{},
				TLSUnique:                   []byte{},
			},
		},
	}
	res, resp, err := client.Bucket.ListMultipartUploads(context.Background(), nil)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(b))
	// b, err = json.MarshalIndent(resp, "", "  ")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	fmt.Println(resp)
}

// 分片上传完成
func TestOk(t *testing.T) {
	client := getCosClient()

	opt := &cos.CompleteMultipartUploadOptions{}
	opt.Parts = append(opt.Parts, cos.Object{
		PartNumber: 1, ETag: "f73ac6d59f0f1228dd44d3139cec8e00"},
	)
	// opt.Parts = append(opt.Parts, cos.Object{
	// 	PartNumber: 2, ETag: "f73ac6d59f0f1228dd44d3139cec8e00"},
	// )
	_, _, err := client.Object.CompleteMultipartUpload(
		context.Background(), Key, UploadId, opt,
	)
	if err != nil {
		panic(err)
	}
}

// 终止分片上传
func TestStop(t *testing.T) {
	client := getCosClient()
	// Abort
	_, err := client.Object.AbortMultipartUpload(context.Background(), Key, UploadId)
	if err != nil {
		panic(err)
	}
}
