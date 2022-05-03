package mutualfund

type MutualFundOption func(*MutualFund)

func NewMutualFund(name string, fundtype MutualFundType, opts ...MutualFundOption) IMutualFund {
	mf := MutualFund{
		Name: name,
	}

	for _, opt := range opts {
		opt(&mf)
	}

	var imf IMutualFund = newOtherMutualFund(mf)

	if fundtype == Mixed {
		imf = newMixedMutualFund(mf)
	}

	if fundtype == Stock {
		imf = newStockMutualFund(mf)
	}

	if fundtype == FixedIncome {
		imf = newFixedIncomeMutualFund(mf)
	}

	if fundtype == MoneyMarket {
		imf = newMoneyMarketMutualFund(mf)
	}

	return imf
}

func WithIndicator(indicator Indicator) MutualFundOption {
	return func(mf *MutualFund) {
		mf.Indicator = indicator
	}
}

func WithAUM(aum float64) MutualFundOption {
	return func(mf *MutualFund) {
		mf.AUM = aum
	}
}
