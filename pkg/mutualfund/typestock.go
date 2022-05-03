package mutualfund

type StockMutualFund struct {
	MutualFund
}

func newStockMutualFund(mf MutualFund) IMutualFund {
	return &StockMutualFund{
		MutualFund: mf,
	}
}
