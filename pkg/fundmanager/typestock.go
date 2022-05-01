package fundmanager

type StockFundManager struct {
	FundManager
}

func newStockFundManager(fm FundManager) IFundManager {
	return &StockFundManager{
		FundManager: fm,
	}
}
