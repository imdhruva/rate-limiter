package fixed_token

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TokenRun struct {
	token     float64
	isSuccess bool
}

var _ = Describe("FixedRate", func() {
	tr1 := []TokenRun{
		{token: 1, isSuccess: true},
		{token: 1, isSuccess: true},
		{token: 1, isSuccess: true},
		{token: 1, isSuccess: true},
		{token: 1, isSuccess: true},
		{token: 1, isSuccess: true},
		{token: 1, isSuccess: false},
		{token: 1, isSuccess: false},
		{token: 1, isSuccess: false},
	}
	Context("validate if requests are allowed", func() {
		It("Should allow request if sent at the rate of 5 with a sleep if .2seconds", func() {
			tb := NewTokenBucket(5, 1)
			for _, v := range tr1 {
				Expect(tb.Request(v.token)).To(Equal(v.isSuccess))
				time.Sleep(time.Millisecond * 200)
			}
		})
	})
})
