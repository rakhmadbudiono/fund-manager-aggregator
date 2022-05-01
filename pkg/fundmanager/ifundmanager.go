package fundmanager

type IFundManager interface {
	GetName() string
	GetIndicator() Indicator
	GetDetail() Detail
}

type FundManager struct {
	Name      string
	Indicator Indicator
	Detail    Detail
}

func (fm *FundManager) GetName() string {
	return fm.Name
}

func (fm *FundManager) GetIndicator() Indicator {
	return fm.Indicator
}

func (fm *FundManager) GetDetail() Detail {
	return fm.Detail
}
