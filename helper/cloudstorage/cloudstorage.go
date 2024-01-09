package cloudstorage

import (
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"path/filepath"
)

var Folder string
var FolderVideo string

func UploadToLocalPath(fileHeader *multipart.FileHeader) (string, error) {
	uploadPath := "./assets/uploads/"
	err := os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	destination := filepath.Join(uploadPath, fileHeader.Filename)
	out, err := os.Create(destination)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		return "", err
	}

	return destination, nil
}

func GetFileName(filePath string) string {
	decodeFilePath, err := url.PathUnescape(filePath)
	if err != nil {
		return ""
	}

	fileName := path.Base(decodeFilePath)
	return fileName
}

func DeleteLocalImage(fileName string) error {
	uploadPath := "./assets/uploads/"

	filePath := filepath.Join(uploadPath, fileName)

	os.Remove(filePath)

	return nil
}
