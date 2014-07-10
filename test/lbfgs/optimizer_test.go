package lbfgs_test

import (
	"github.com/igorbonadio/optimus/lbfgs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Optimizer", func() {
	It("should find the minimum of f1(x) = x1^2 + x1 + x2^2", func() {
		opt := lbfgs.New(
			func(x []float64) float64 { return x[0]*x[0] + x[0] + x[1]*x[1] },
			func(x []float64) []float64 { return []float64{2*x[0] + 1, 2 * x[1]} },
			5, 20, 0.0001)
		min := opt.Min([]float64{1, 1})
		Ω(min[0]).Should(BeNumerically("~", -0.5, 0.001))
		Ω(min[1]).Should(BeNumerically("~", 0, 0.001))
	})

	It("should find the minimum of f1(x) = x1^4 + x1 + x2^2", func() {
		opt := lbfgs.New(
			func(x []float64) float64 { return x[0]*x[0]*x[0]*x[0] + x[0] + x[1]*x[1] },
			func(x []float64) []float64 { return []float64{4*x[0]*x[0]*x[0] + 1, 2 * x[1]} },
			5, 20, 0.0001)
		min := opt.Min([]float64{1, 1})
		Ω(min[0]).Should(BeNumerically("~", -0.6299, 0.001))
		Ω(min[1]).Should(BeNumerically("~", 0, 0.001))
	})
})
