package heif

import (
	"fmt"
	"os"

	"go4.org/media/heif"
)

func OpenAndGetExifBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	heifFile := heif.Open(file)

	exif, err := heifFile.EXIF()
	if err != nil {
		return nil, fmt.Errorf("heifFile.EXIF: %w", err)
	}

	return exif, nil
}
