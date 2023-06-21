package main

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestShadowing(t *testing.T) {
	for _, c := range []bool{true, false} {
		t.Run(fmt.Sprint("case where tracing is ", c), func(t *testing.T) {
			g := NewWithT(t)

			g.Ω(Shadowing(c)).To(BeNil()) // Client and error are nil values
		})
	}
}

func TestShadowingPossibleSolution(t *testing.T) {
	for _, c := range []bool{true, false} {
		t.Run(fmt.Sprint("case where tracing is ", c), func(t *testing.T) {
			g := NewWithT(t)

			g.Ω(ShadowingPossibleSolution(c)).NotTo(BeNil()) // Client and error are nil values
		})
	}
}

func TestShadowingReassignmentSolution(t *testing.T) {
	for _, c := range []bool{true, false} {
		t.Run(fmt.Sprint("case where tracing is ", c), func(t *testing.T) {
			g := NewWithT(t)

			g.Ω(ShadowingReassignmentSolution(c)).NotTo(BeNil()) // Client and error are nil values
		})
	}
}

func TestShadowingPossibleSolutionOther(t *testing.T) {
	for _, c := range []bool{true, false} {
		t.Run(fmt.Sprint("case where tracing is ", c), func(t *testing.T) {
			g := NewWithT(t)

			g.Ω(ShadowingPossibleSolutionOther(c)).NotTo(BeNil()) // Client and error are nil values
		})
	}
}

func TestShadowingPossibleSolutionOther2(t *testing.T) {
	for _, c := range []bool{true, false} {
		t.Run(fmt.Sprint("case where tracing is ", c), func(t *testing.T) {
			g := NewWithT(t)

			g.Ω(ShadowingPossibleSolutionOther2(c)).NotTo(BeNil()) // Client and error are nil values
		})
	}
}
