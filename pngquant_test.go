package pngquant

import (
	"bytes"
	"image/png"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	var file *os.File
	file, _ = os.Open("gopher.png")
	defer file.Close()
	info, _ := file.Stat()
	orgSize := info.Size()
	orgImg, _ := png.Decode(file)

	// Test with good data
	newImg, err := Compress(orgImg, "1")
	if err != nil {
		t.Errorf("error has occurred: %v", err)
	} else {
		var w bytes.Buffer
		png.Encode(&w, newImg)
		if len(w.Bytes()) > int(orgSize) {
			t.Error("image is not comppressed")
		}
	}

	// Test with non-numerical data (no need to test encode as we expect nothing)
	_, err = Compress(orgImg, "one")
	if err == nil {
		t.Errorf("ERROR: Expected failure on incompatible speed argument (one)")
	}

	// Test with number that exceeds the limit (no need to test encode as we expect nothing)
	_, err = Compress(orgImg, "20")
	if err == nil {
		t.Errorf("ERROR: Expected failure on incompatible speed argument (20)")
	}
}
