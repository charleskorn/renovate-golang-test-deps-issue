package main_test

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestThing(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect("hello").To(Equal("hello"))
}
