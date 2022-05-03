package mutualfund

type IMutualFund interface {
	GetName() string
	GetIndicator() Indicator
	GetDetail() Detail
}

type MutualFund struct {
	Name      string
	Indicator Indicator
	Detail    Detail
}

func (mf *MutualFund) GetName() string {
	return mf.Name
}

func (mf *MutualFund) GetIndicator() Indicator {
	return mf.Indicator
}

func (mf *MutualFund) GetDetail() Detail {
	return mf.Detail
}
