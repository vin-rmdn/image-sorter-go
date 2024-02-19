package heif_test

import (
	"testing"

	localheif "image-sorter/pkg/image/heif"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpenAndGetExifBytes(t *testing.T) {
	t.Run("when os open fails", func(t *testing.T) {
		actual, err := localheif.OpenAndGetExifBytes("testdata/sample.heic")

		require.NoError(t, err)
		assert.Contains(t, string(actual), "Exif")
	})
}
