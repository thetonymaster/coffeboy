package imageutil_test

import (
	"io/ioutil"

	. "github.com/crowdint/coffeboy/imageutil"

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
})
