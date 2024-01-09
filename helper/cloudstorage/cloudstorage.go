package cloudstorage

import (
	"context"
	"io"
	"mime/multipart"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

var Folder string
var FolderVideo string

func UploadToLocalPath(fileHeader *multipart.FileHeader) (string, error) {
	uploadPath := "./assets/upload/"
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

func DeleteImage(fileName string) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("storage.json"))
	if err != nil {
		return echo.NewHTTPError(500, err)
	}

	bucketName := "ecowave"
	objectPath := Folder + fileName

	obj := client.Bucket(bucketName).Object(objectPath)

	err = obj.Delete(ctx)
	if err != nil {
		return echo.NewHTTPError(500, "Gagal menghapus file pada cloud storage")
	}

	return nil
}
