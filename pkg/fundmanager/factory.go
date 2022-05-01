package fundmanager

const (
	FixedIncome = "Pendapatan Tetap"
	Mixed       = "Campuran"
	MoneyMarket = "Pasar Uang"
	Stock       = "Saham"
)

type FundManagerOption func(*FundManager)

func NewFundManager(name, fundtype string, opts ...FundManagerOption) IFundManager {
	fm := FundManager{
		Name: name,
	}

	for _, opt := range opts {
		opt(&fm)
	}

	var ifm IFundManager = newOtherFundManager(fm)

	if fundtype == FixedIncome {
		ifm = newFixedIncomeFundManager(fm)
	}

	if fundtype == Mixed {
		ifm = newMixedFundManager(fm)
	}

	if fundtype == MoneyMarket {
		ifm = newMoneyMarketFundManager(fm)
	}

	if fundtype == Stock {
		ifm = newStockFundManager(fm)
	}

	return ifm
}

func WithIndicator(indicator Indicator) FundManagerOption {
	return func(fm *FundManager) {
		fm.Indicator = indicator
	}
}

func WithDetail(detail Detail) FundManagerOption {
	return func(fm *FundManager) {
		fm.Detail = detail
	}
}
