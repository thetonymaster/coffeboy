package utils_test

import (
	"io/ioutil"

	. "github.com/crowdint/coffeboy/utils"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Utils", func() {
	Describe("Upload and resize a file", func() {
		Context("From the hard drive", func() {

			auth, err := aws.EnvAuth()
			if err != nil {
				panic(err)
			}

			ss := s3.New(auth, aws.Regions["us-west-2"])

			bucket := ss.Bucket("coffeboy")

			fileBytes, err := ioutil.ReadFile("../testfiles/test.jpg")
			if err != nil {
				panic(err)
			}

			It("Should manage all the tasks", func() {
				UploadAndResizeImage(100, 100, fileBytes, "test.jpg")

				bucketFile, err := bucket.Get("test.jpg")
				Expect(err).To(BeNil())
				Expect(len(bucketFile)).Should(BeNumerically("<", len(fileBytes)))

				err = bucket.Del("test.jpg")
				Expect(err).To(BeNil())

			})
		})
	})
})
