package hello_test

import (
	"testing"

	"github.com/ivanfetch/go-hello-world"
)

func TestHello(t *testing.T) {
	t.Parallel()
	// This doesn't really test anything, but does import our hello package,
	// which is the point in this case.
	hello.Print()
}
