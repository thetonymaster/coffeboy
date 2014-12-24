package utils

import "github.com/crowdint/coffeboy/imageutil"

func UploadAndResizeImage(maxWidth, maxHeight uint, fileBytes []byte, path string) error {
	fileResized, err := imageutil.Resize(fileBytes, 100, 100)
	if err != nil {
		return err
	}

	err = imageutil.Upload(fileResized, path)
	if err != nil {
		return err
	}

	return nil
}
