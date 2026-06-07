package tempfile_test

import (
	"testing"

	"github.com/kopia/kopia/infernal/tempfile"
)

func TestTempFile(t *testing.T) {
	tempfile.VerifyTempfile(t, tempfile.CreateAutoDelete)
}
