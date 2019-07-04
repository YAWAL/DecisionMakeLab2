package calculation

const years = 5

type Build struct {
	Name              string
	Value             float64
	Profit            float64
	ProfitProbability float64
	Loss              float64
	LossProbability   float64
}

type Wait struct {
	Name                 string
	PositiveProbability  float64
	NegativeProbability  float64
	NewBuildA, NewBuildB Build
}

func (b *Build) ProcessBuild() float64 {
	return years*(b.Profit*b.ProfitProbability+b.Loss*b.LossProbability) - b.Value
}

func (w *Wait) ProcessWait() float64 {
	resultA1 := w.NewBuildA.ProcessBuild()
	resultB1 := w.NewBuildB.ProcessBuild()
	if resultA1 > resultB1 {
		return w.PositiveProbability * resultA1
	}
	return w.PositiveProbability * resultB1
}
