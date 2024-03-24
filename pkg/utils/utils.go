package utils

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/nfnt/resize"
)

const (
	basePath  = "./assets/"
	imagePath = "/assets"
)

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func SaveFile(file *multipart.FileHeader) (*string, error) {
	var img image.Image
	var imageExt string
	if file == nil {
		return nil, nil
	}

	arr := strings.Split(file.Filename, ".")
	if len(arr) > 1 {
		imageExt = arr[len(arr)-1]
	}
	newFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer newFile.Close()

	switch imageExt {
	case "jpeg":
		img, err = jpeg.Decode(newFile)
	case "jpg":
		img, err = jpeg.Decode(newFile)
	case "png":
		img, err = png.Decode(newFile)
	case "PNG":
		img, err = png.Decode(newFile)
	case "gif":
		img, err = gif.Decode(newFile)
	default:
		err = errors.New("Unsupported file type")
		return nil, err
	}
	fullFileName := fmt.Sprintf("%s.%s", uuid.NewString(), imageExt)

	fileOnDisk, err := os.Create(fmt.Sprintf("%s/%s", basePath, fullFileName))

	defer fileOnDisk.Close()

	if err != nil {
		return nil, err
	}

	imageWidth := uint(MinInt(1000, img.Bounds().Max.X))

	resizedImg := resize.Resize(imageWidth, 0, img, resize.Lanczos3)

	err = jpeg.Encode(fileOnDisk, resizedImg, nil)

	if err != nil {
		return nil, err
	}
	path := fmt.Sprintf("%s/%s", imagePath, fullFileName)
	return &path, err
}
