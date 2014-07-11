package lbfgs

import (
	"github.com/igorbonadio/lbfgs"
)

type Optimizer struct {
	f            ObjectiveFunction
	df           PartialDerivatives
	m            int
	maxIteration int
	epsilon      float64
}

type ObjectiveFunction func(x []float64) float64
type PartialDerivatives func(x []float64) []float64

func New(f ObjectiveFunction, df PartialDerivatives, m int, maxIteration int, epsilon float64) *Optimizer {
	return &Optimizer{f: f, df: df, m: m, maxIteration: maxIteration, epsilon: epsilon}
}

func (optimizer *Optimizer) Min(initialX []float64) []float64 {
	opt := lbfgs.NewOptimizer(optimizer.m, len(initialX))

	optX := initialX

	x := lbfgs.NewVector(len(initialX))
	g := lbfgs.NewVector(len(initialX))
	x.SetValues(optX)

	optimizer.f(optX)

	for k := 0; k < optimizer.maxIteration; k++ {
		g.SetValues(optimizer.df(optX))
		delta := opt.GetDeltaX(x, g)
		x.Increment(delta, 1)
		for j := 0; j < len(optX); j++ {
			optX[j] = x.Get(j)
		}
		optimizer.f(optX)
		if g.Norm() < optimizer.epsilon /*|| math.Abs((fx-newFx)/newFx) < optimizer.epsilon*/ {
			break
		}
	}

	return optX
}
