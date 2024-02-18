package fixed_token

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestFixedRate(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fixed Rate Suite")
}
