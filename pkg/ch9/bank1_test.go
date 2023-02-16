package ch9

import (
	"github.com/AthulMuralidhar/the-go-book/pkg/ch9/bank1"
	"testing"
)

func Test_teller(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bank1.teller()
		})
	}
}
