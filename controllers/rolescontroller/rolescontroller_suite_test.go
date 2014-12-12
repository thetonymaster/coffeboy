package rolescontroller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRolescontroller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Rolescontroller Suite")
}
