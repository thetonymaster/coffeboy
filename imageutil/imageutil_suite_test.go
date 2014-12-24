package imageutil_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestImageutil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Imageutil Suite")
}
