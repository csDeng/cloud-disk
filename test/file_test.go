package test

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path"
	"strconv"
	"testing"
)

const (
	// 分片大小 100M
	chunkSize = 1 * 1024 * 1024
)

// 文件分片
func TestGenChunkFile(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal("get pwd fail, error: ", err)
	}
	name := path.Join(pwd, "/big_file/os.pdf")

	// 0666 => _rw_ rw_ rw_
	// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.

	fileInfo, err := os.Stat(name)
	if err != nil {
		t.Fatal(err)
	}
	myfile, err := os.OpenFile(name, os.O_RDONLY, 0666)

	if err != nil {
		t.Fatal(err)
	}
	defer myfile.Close()
	// 分片个数
	// int(80 / 30) = 2
	// chunkNum := int(fileInfo.Size() / chunkSize)
	// 向上取整
	chunkNum := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))
	log.Printf("chunkNum = %d", chunkNum)
	b := make([]byte, chunkSize)

	var rest int64
	for i := 0; i < chunkNum; i++ {
		// 指定读取文件的起始位置
		myfile.Seek(int64(i*chunkSize), 0)

		// 剩下的未读字节
		rest = fileInfo.Size() - int64(i*chunkSize)
		if chunkSize > rest {
			b = make([]byte, rest)
		}
		// 读取数据
		n, err := myfile.Read(b)
		if err != nil {
			t.Fatalf("读取文件出现错误, rr = %v", err.Error())
		}
		log.Printf("第 %d 轮, 读取了 %d byte \r\n", i, n)
		f, err := os.OpenFile("./"+strconv.Itoa(i)+"_chunk", os.O_CREATE|os.O_WRONLY, os.ModePerm)
		if err != nil {
			t.Fatalf("分片存储, err = %s", err.Error())
		}
		f.Write(b)
		f.Close()
	}

}

// 文件合并
func TestMerge(t *testing.T) {
	myfile, err := os.OpenFile("./test.pdf", os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}
	defer myfile.Close()
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal("get pwd fail, error: ", err)
	}
	name := path.Join(pwd, "/big_file/os.pdf")

	// 0666 => _rw_ rw_ rw_
	// Stat returns a FileInfo describing the named file. If there is an error, it will be of type *PathError.

	fileInfo, err := os.Stat(name)
	if err != nil {
		t.Fatal(err)
	}

	chunkNum := int(math.Ceil(float64(fileInfo.Size()) / float64(chunkSize)))
	for i := 0; i < chunkNum; i++ {
		f, err := os.OpenFile("./"+strconv.Itoa(i)+"_chunk", os.O_RDONLY, os.ModePerm)
		if err != nil {
			t.Fatalf("读取分片, err = %s", err.Error())
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			t.Fatal(err)
		}
		myfile.Write(b)
		f.Close()
	}

}

// 文件一致性校验
func TestCheck(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal("get pwd fail, error: ", err)
	}
	name := path.Join(pwd, "/big_file/os.pdf")
	f1, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f1.Close()
	b1, err := ioutil.ReadAll(f1)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.OpenFile("./test.pdf", os.O_RDONLY, 0666)
	if err != nil {
		t.Fatal(err)
	}
	defer f2.Close()
	b2, err := ioutil.ReadAll(f2)
	if err != nil {
		t.Fatal(err)
	}

	// 2^8 -> 2^4 => 16bit => 32bit
	s1 := fmt.Sprintf("%x", md5.Sum(b1))
	s2 := fmt.Sprintf("%x", md5.Sum(b2))
	fmt.Println(s1, s2, s1 == s2)
}
