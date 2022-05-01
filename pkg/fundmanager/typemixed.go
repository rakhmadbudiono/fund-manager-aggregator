package fundmanager

type MixedFundManager struct {
	FundManager
}

func newMixedFundManager(fm FundManager) IFundManager {
	return &MixedFundManager{
		FundManager: fm,
	}
}
