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
	"github.com/satori/go.uuid"
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
		fmt.Println(err)
		return nil
	}
	return c
}

func uploadToStorage(fileData []byte, fileName string) (string, error) {
	bucketName := config.BucketName()

	ctx := context.Background()

	fmt.Println(bucketName)

	name := uuid.NewV4().String() + fileName
	w := c.Bucket(bucketName).Object(name).NewWriter(ctx)
	w.ACL = []storage.ACLRule{{Entity: storage.AllUsers, Role: storage.RoleReader}}

	defer w.Close()

	if n, err := w.Write(fileData); err != nil {
		return "", err
	} else if n != len(fileData) {
		return "", err
	}
	if err := w.Close(); err != nil {
		return "", err
	}

	publicURL := "https://storage.googleapis.com/" + bucketName + "/" + name

	return publicURL, nil

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
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(len(buf))
	return buf, nil
}

func UploadFile(localFilePath, remoteFileName string) (string, error) {
	buf, err := readFileBinary(localFilePath)
	url, err := uploadToStorage(buf, remoteFileName)
	return url, err

}

func UploadBinaryData(data []byte, fileName string) (string, error) {
	url, err := uploadToStorage(data, fileName)
	return url, err
}
