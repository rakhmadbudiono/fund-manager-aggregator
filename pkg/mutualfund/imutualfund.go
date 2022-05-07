package mutualfund

import (
	"fmt"
)

type MutualFundType int

const (
	Mixed MutualFundType = iota
	Stock
	FixedIncome
	MoneyMarket
)

type IMutualFund interface {
	GetName() string
	GetIndicator() Indicator
	GetAUM() float64
	ToArrayString() []string
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

func (mf *MutualFund) ToArrayString() []string {
	arr := []string{}

	arr = append(arr, mf.Name)
	arr = append(arr, mf.Indicator.toArrayString()...)
	arr = append(arr, fmt.Sprintf("%f", mf.AUM))

	return arr
}
