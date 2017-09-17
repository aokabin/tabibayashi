package storage

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/aokabin/tabibayashi/config"
)

var (
	c *storage.Client
)

func init() {
	rand.Seed(time.Now().UnixNano())
	c = Connection()
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString1(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func Connection() *storage.Client {
	ctx := context.Background()
	c, err := storage.NewClient(ctx)
	if err != nil {
		return nil
	}
	return c
}

func uploadToStorage(fileData []byte, fileName string) error {
	bucketName := config.BucketName()

	ctx := context.Background()

	fmt.Println(bucketName)

	w := c.Bucket(bucketName).Object(fileName).NewWriter(ctx)
	defer w.Close()

	if n, err := w.Write(fileData); err != nil {
		return err
	} else if n != len(fileData) {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	return nil

}

func readFileBinary(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(buf))
	return buf, nil
}

func UploadFile(localFilePath, remoteFileName string) error {
	buf, err := readFileBinary(localFilePath)
	err = uploadToStorage(buf, remoteFileName)
	return err

}

func UploadBinaryData(data []byte, fileName string) error {
	return uploadToStorage(data, fileName)
}
