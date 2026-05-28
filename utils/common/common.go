package common

import (
	"encoding/base64"
	"image"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

func DecodeBase64(base64String string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64String)
}

func SaveAsBMP(binaryData []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(strings.NewReader(string(binaryData)))
	if err != nil {
		return err
	}

	return bmp.Encode(file, img)
}
