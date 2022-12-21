package test

import "testing"
import gomega "github.com/onsi/gomega"

func TestFarmHasCow(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	//f := farm.New([]string{"Cow", "Horse"})
	g.Expect(0).To(gomega.Equal(0))
}
