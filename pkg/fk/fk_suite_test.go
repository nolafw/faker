package fk_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFk(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fk Suite")
}
