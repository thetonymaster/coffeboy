package roles_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRoles(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roles Suite")
}
