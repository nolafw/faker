package internet_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestInternet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Internet Suite")
}
