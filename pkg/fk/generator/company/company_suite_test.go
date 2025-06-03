package company_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCompany(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Company Suite")
}
