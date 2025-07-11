package faker_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFaker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Faker Suite")
}
