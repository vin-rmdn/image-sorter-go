package exif_test

import (
	"testing"
	"time"

	"image-sorter/pkg/image/exif"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTimeFromEXIF(t *testing.T) {
	t.Run("errors", func(t *testing.T) {
		t.Run("when file does not exist", func(t *testing.T) {
			actual, err := exif.TimeFromFile("testdata/file does not exist")
			require.EqualError(t, err, "os.Open: open testdata/file does not exist: no such file or directory")
			assert.Empty(t, actual)
		})

		t.Run("when image has no exif", func(t *testing.T) {
			actual, err := exif.TimeFromFile("testdata/sample without exif.jpg")
			require.EqualError(t, err, "exif.SearchAndExtractExif: no exif data")
			assert.Empty(t, actual)
		})

		// TODO: do corrupt exif!
	})

	t.Run("successful", func(t *testing.T) {
		for testName, testFileName := range map[string]string{
			"heif": "testdata/sample.heic",
			"jpeg": "testdata/sample.jpg",
		} {
			t.Run(testName, func(t *testing.T) {
				actual, err := exif.TimeFromFile(testFileName)
				require.NoError(t, err)

				assert.Equal(t,
					time.Date(2024, 2, 19, 5, 41, 29, 0, time.UTC),
					actual.UTC(),
				)
			})
		}
	})
}
