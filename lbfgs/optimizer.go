package lbfgs

import (
	lbfgs "github.com/kho/liblbfgs/go"
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

func (optimizer *Optimizer) Min(x []float64) []float64 {
	evaluate := func(x, g []lbfgs.Float, step lbfgs.Float) lbfgs.Float {
		_x := make([]float64, len(x))
		for i := 0; i < len(x); i++ {
			_x[i] = float64(x[i])
		}
		_g := optimizer.df(_x)
		for i := 0; i < len(_g); i++ {
			g[i] = lbfgs.Float(_g[i])
		}
		return lbfgs.Float(optimizer.f(_x))
	}

	_x := make([]lbfgs.Float, len(x))
	for i := 0; i < len(x); i++ {
		_x[i] = lbfgs.Float(x[i])
	}

	lbfgs.Minimize(_x, evaluate, lbfgs.Silent, &lbfgs.DefaultParam)

	for i := 0; i < len(x); i++ {
		x[i] = float64(_x[i])
	}

	return x
}
