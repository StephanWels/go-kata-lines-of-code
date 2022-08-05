package loc

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLinesOfCode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CountLinesOfJavaCode Test Suite")
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}

var _ = Describe("java line of code counter", func() {
	It("should ignore simple comments", func() {
		dat, err := os.ReadFile("test_programs/test_program_2.java")
		check(err)
		Expect(CountLinesOfJavaCode(string(dat))).To(Equal(2))
	})
})
