package imageutil_test

import (
	"io/ioutil"

	. "github.com/crowdint/coffeboy/imageutil"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Imageutil", func() {
	Describe("Create a thumbnail from an image", func() {
		Context("Get an image byte array from hard drive", func() {

			fileBytes, err := ioutil.ReadFile("../testfiles/test.jpg")
			if err != nil {
				panic(err)
			}

			It("Should return an array of bytes and no error", func() {
				fileResized, err := Resize(fileBytes, 100, 100)
				Expect(err).To(BeNil())
				Expect(len(fileResized)).Should(BeNumerically("<", len(fileBytes)))

			})
		})
	})

	Describe("Upload a file to S3", func() {
		Context("With a test file from the drive", func() {

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

			It("Should Upload a file and return no error", func() {
				err := Upload(fileBytes, "test.jpg")
				Expect(err).To(BeNil())

				bucketFile, err := bucket.Get("test.jpg")

				Expect(bucketFile).Should(Equal(fileBytes))

				err = bucket.Del("test.jpg")
				Expect(err).To(BeNil())

			})
		})
	})
})
