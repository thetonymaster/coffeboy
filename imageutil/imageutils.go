package imageutil

import (
	"bytes"
	"image/jpeg"

	"github.com/nfnt/resize"
)

func Resize(fileBytes []byte, maxWidth, maxHeight uint) ([]byte, error) {
	reader := bytes.NewReader(fileBytes)

	img, err := jpeg.Decode(reader)
	if err != nil {
		return nil, err
	}

	m := resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)
	out := bytes.NewBuffer([]byte{})

	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return nil, err
	}
	outBytes := out.Bytes()
	return outBytes, nil

}
