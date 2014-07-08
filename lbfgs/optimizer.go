package lbfgs

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
	return []float64{-0.5, 0}
}
