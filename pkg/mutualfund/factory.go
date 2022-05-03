package mutualfund

const (
	FixedIncome = "Pendapatan Tetap"
	Mixed       = "Campuran"
	MoneyMarket = "Pasar Uang"
	Stock       = "Saham"
)

type MutualFundOption func(*MutualFund)

func NewMutualFund(name, fundtype string, opts ...MutualFundOption) IMutualFund {
	mf := MutualFund{
		Name: name,
	}

	for _, opt := range opts {
		opt(&mf)
	}

	var imf IMutualFund = newOtherMutualFund(mf)

	if fundtype == FixedIncome {
		imf = newFixedIncomeMutualFund(mf)
	}

	if fundtype == Mixed {
		imf = newMixedMutualFund(mf)
	}

	if fundtype == MoneyMarket {
		imf = newMoneyMarketMutualFund(mf)
	}

	if fundtype == Stock {
		imf = newStockMutualFund(mf)
	}

	return imf
}

func WithIndicator(indicator Indicator) MutualFundOption {
	return func(mf *MutualFund) {
		mf.Indicator = indicator
	}
}

func WithDetail(detail Detail) MutualFundOption {
	return func(mf *MutualFund) {
		mf.Detail = detail
	}
}
