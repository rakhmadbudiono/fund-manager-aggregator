package fundmanager

type FixedIncomeFundManager struct {
	FundManager
}

func newFixedIncomeFundManager(fm FundManager) IFundManager {
	return &FixedIncomeFundManager{
		FundManager: fm,
	}
}
