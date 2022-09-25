package file_helper

import (
	"cloud_disk/app/common/helper"
	"log"
	"os"
)

func CreateFile(data []byte, ext string) (string, error) {

	p := "../pool/" + helper.GenerateUuid() + ext
	f, err := os.Create(p)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = f.Write(data)
	if nil != err {
		return "", err
	}
	log.Println("create file=>", p)
	return p, nil
}
