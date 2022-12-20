package test

import (
	"fmt"
	"github.com/AthulMuralidhar/the-go-book/pkg/ch1"
	"os"
	"testing"
)

func BenchmarkEcho2(b *testing.B) {
	err := os.Setenv("GO_ENV", ch1.TestingEnv)
	if err != nil {
		_ = fmt.Errorf("error during TestMain: %w\n", err)
		return
	}

	for i := 0; i < b.N; i++ {
		ch1.Echo2()
	}
}
