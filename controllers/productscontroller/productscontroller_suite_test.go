package productscontroller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestProductscontroller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Productscontroller Suite")
}
