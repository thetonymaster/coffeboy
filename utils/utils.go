package utils

import (
	"log"

	"github.com/crowdint/coffeboy/imageutil"
)

func UploadAndResizeImage(maxWidth, maxHeight uint, fileBytes []byte, path string) {
	fileResized, err := imageutil.Resize(fileBytes, maxWidth, maxHeight)
	if err != nil {
		log.Printf("Error %s\n", err.Error())
		return
	}

	err = imageutil.Upload(fileResized, path)
	if err != nil {
		log.Printf("Error %s\n", err.Error())
		return
	}

}

func UploadImages(fileBytes []byte, identifier, folder string) {
	go imageutil.Upload(fileBytes, folder+"/"+identifier+".jpg")
	go UploadAndResizeImage(400, 300, fileBytes, folder+"/"+identifier+"_small.jpg")
}
