package fundmanager

type MoneyMarketFundManager struct {
	FundManager
}

func newMoneyMarketFundManager(fm FundManager) IFundManager {
	return &MoneyMarketFundManager{
		FundManager: fm,
	}
}
