package lbfgs

import (
	"github.com/huichen/lbfgs"
	"math"
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

	optX := SF64ToSF32(initialX)

	x := lbfgs.NewVector(len(initialX))
	g := lbfgs.NewVector(len(initialX))
	x.SetValues(optX)

	fx := optimizer.f(SF32ToSF64(optX))

	for k := 0; k < optimizer.maxIteration; k++ {
		g.SetValues(SF64ToSF32(optimizer.df(SF32ToSF64(optX))))
		delta := opt.GetDeltaX(x, g)
		x.Increment(delta, 1)
		for j := 0; j < len(optX); j++ {
			optX[j] = x.Get(j)
		}
		newFx := optimizer.f(SF32ToSF64(optX))
		if g.Norm() < float32(optimizer.epsilon) || math.Abs((fx-newFx)/newFx) < optimizer.epsilon {
			break
		}
	}

	return SF32ToSF64(optX)
}

func SF64ToSF32(x []float64) []float32 {
	_x := make([]float32, len(x))
	for i := 0; i < len(x); i++ {
		_x[i] = float32(x[i])
	}
	return _x
}

func SF32ToSF64(x []float32) []float64 {
	_x := make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		_x[i] = float64(x[i])
	}
	return _x
}
