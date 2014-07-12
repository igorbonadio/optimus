package lbfgs_test

import (
	. "github.com/kho/liblbfgs/go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func evaluate(x, g []Float, step Float) Float {
	g[0] = 2*x[0] + 1
	g[1] = 2 * x[1]
	return x[0]*x[0] + x[0] + x[1]*x[1]
}

var _ = Describe("LibLBFGS", func() {
	It("should find the minimum of f1(x) = x1^2 + x1 + x2^2", func() {
		x := []Float{1, 1}
		Minimize(x, evaluate, Silent, &DefaultParam)
		Ω(x[0]).Should(BeNumerically("~", -0.5, 0.001))
		Ω(x[1]).Should(BeNumerically("~", 0, 0.001))
	})
})
