package userscontroller_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUserscontroller(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Userscontroller Suite")
}
