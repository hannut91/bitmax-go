package bitmax

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBitmaxGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BitmaxGo Suite")
}
