package currenttime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCurrenttime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Currenttime Suite")
}
