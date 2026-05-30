package common

import (
	"encoding/base64"
	"fmt"
	"image"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/image/bmp"
)

func DecodeBase64(base64String string) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 string: %w", err)
	}
	return data, nil
}

func SaveAsBMP(binaryData []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer SafeClose(file)

	img, _, err := image.Decode(strings.NewReader(string(binaryData)))
	if err != nil {
		return fmt.Errorf("error decode image: %w", err)
	}
	if err := bmp.Encode(file, img); err != nil {
		return fmt.Errorf("encode image to bmp: %w", err)
	}
	return nil
}

func SafeClose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("close failed: %v", err)
	}
}
