package mutualfund

type IMutualFund interface {
	GetName() string
	GetIndicator() Indicator
	GetAUM() float64
}

type MutualFund struct {
	Name      string
	Indicator Indicator
	AUM       float64
}

func (mf *MutualFund) GetName() string {
	return mf.Name
}

func (mf *MutualFund) GetIndicator() Indicator {
	return mf.Indicator
}

func (mf *MutualFund) GetAUM() float64 {
	return mf.AUM
}
