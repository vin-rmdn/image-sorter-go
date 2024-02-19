package exif

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dsoprea/go-exif/v3"
)

const (
	idDateTimeOriginal = 0x9003
	idOffsetTime       = 0x9010
)

func TimeFromFile(filePath string) (time.Time, error) {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return time.Time{}, fmt.Errorf("os.Open: %w", err)
	}

	exifData, err := exif.SearchAndExtractExif(fileData)
	if err != nil {
		return time.Time{}, fmt.Errorf("exif.SearchAndExtractExif: %w", err)
	}

	entries, _, err := exif.GetFlatExifData(exifData, nil)
	if err != nil {
		return time.Time{}, fmt.Errorf("exif.GetFlatExifData: %w", err)
	}

	var dateTimeString string
	var offsetTime string
	var ok bool
	for _, entry := range entries {
		if entry.TagId == idDateTimeOriginal {
			dateTimeString, ok = entry.Value.(string)
			if !ok {
				return time.Time{}, errors.New("dateTimeString, ok := entry.Value.(string)")
			}
		} else if entry.TagId == idOffsetTime {
			offsetTime, ok = entry.Value.(string)
			if !ok {
				return time.Time{}, errors.New("offsetTime, ok := entry.Value.(string)")
			}
		}
	}

	if len(dateTimeString) == 0 || len(offsetTime) == 0 {
		return time.Time{}, errors.New("date time or offset if empty")
	}

	exifTime, err := time.Parse(
		"2006:01:02 15:04:05 -07:00",
		fmt.Sprintf("%s %s", dateTimeString, offsetTime),
	)
	if err != nil {
		return time.Time{}, fmt.Errorf("time.Parse: %w", err)
	}

	return exifTime, nil
}
