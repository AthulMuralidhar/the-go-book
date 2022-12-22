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
	temp1, err := setupTempfile()
	defer func(temp1 *os.File) {
		err := temp1.Close()
		if err != nil {
			log.Panic(err)
		}
	}(temp1)
	if err != nil {
		log.Panic(err)
	}
	temp2, err := setupTempfile()
	defer func(temp2 *os.File) {
		err := temp2.Close()
		if err != nil {
			log.Panic(err)
		}
	}(temp2)
	if err != nil {
		log.Panic(err)
	}

}
func setupTempfile() (*os.File, error) {
	content := []byte("temporary file's content")
	tempfile, err := ioutil.TempFile("./", "example.*.txt")
	if err != nil {
		return nil, err
	}

	if _, err := tempfile.Write(content); err != nil {
		err := tempfile.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}
	if err := tempfile.Close(); err != nil {
		return nil, err
	}
	return tempfile, nil
}
