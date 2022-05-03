package mutualfund

type MoneyMarketMutualFund struct {
	MutualFund
}

func newMoneyMarketMutualFund(mf MutualFund) IMutualFund {
	return &MoneyMarketMutualFund{
		MutualFund: mf,
	}
}
