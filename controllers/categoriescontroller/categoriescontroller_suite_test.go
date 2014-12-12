package categoriescontroller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCategoriescontroller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Categoriescontroller Suite")
}
