package test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)
import gomega "github.com/onsi/gomega"

func TestEx1_4(t *testing.T) {
	g := gomega.NewGomegaWithT(t)

	//f := farm.New([]string{"Cow", "Horse"})
	//g.Expect(0).To(gomega.Equal(0))

}
func setup() {
	content1 := []byte("temporary file's content")
	tmpfile1, err := ioutil.TempFile("./", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(name string) {
		err := os.Remove(name)
		if err != nil {

		}
	}(tmpfile1.Name()) // clean up

	if _, err := tmpfile1.Write(content1); err != nil {
		err := tmpfile1.Close()
		if err != nil {
			return
		}
		log.Fatal(err)
	}
	if err := tmpfile1.Close(); err != nil {
		log.Fatal(err)
	}
}
