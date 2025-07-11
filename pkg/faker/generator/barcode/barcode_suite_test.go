package barcode_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBarcode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Barcode Suite")
}
