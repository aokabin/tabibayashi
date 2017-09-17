package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadMusicFile(t *testing.T) {
	err := UploadFile("../README.md", "sample")
	assert.Nil(t, err)
}

func TestUploadBinaryData(t *testing.T) {
	buf, err := readFileBinary("../README.md")
	assert.Nil(t, err)
	err = UploadBinaryData(buf, "yonple")
	assert.Nil(t, err)
}
