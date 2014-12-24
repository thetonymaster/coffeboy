package imageutil

import (
	"bytes"
	"image/jpeg"

	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
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

func Upload(fileBytes []byte, path string) error {
	auth, err := aws.EnvAuth()
	if err != nil {
		return err
	}

	ss := s3.New(auth, aws.Regions["us-west-2"])

	bucket := ss.Bucket("coffeboy")

	err = bucket.Put(path, fileBytes, "image/jpeg", s3.PublicRead, s3.Options{})

	return err

}
