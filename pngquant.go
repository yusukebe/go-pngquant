package pngquant

import (
	"bytes"
	"fmt"
	"image"
	"image/png"
	"os/exec"
	"strconv"
	"strings"
)

func Compress(input image.Image, speed string) (output image.Image, err error) {
	if err = speedCheck(speed); err != nil {
		return
	}

	var w bytes.Buffer
	err = png.Encode(&w, input)
	if err != nil {
		return
	}

	b := w.Bytes()
	compressed, err := CompressBytes(b, speed)
	if err != nil {
		return
	}

	output, err = png.Decode(bytes.NewReader(compressed))
	return
}

func CompressBytes(input []byte, speed string) (output []byte, err error) {
	cmd := exec.Command("pngquant", "-", "--speed", speed)
	cmd.Stdin = strings.NewReader(string(input))
	var o bytes.Buffer
	cmd.Stdout = &o
	err = cmd.Run()

	if err != nil {
		return
	}

	output = o.Bytes()
	return
}

func speedCheck(speed string) (err error) {
	// conversion, as an aside, also forces the speed argument to be a number.
	speedInt, err := strconv.Atoi(speed)
	if err != nil {
		return
	}

	if speedInt > 10 {
		return fmt.Errorf("speed cannot exceed value of 10")
	}

	return
}
