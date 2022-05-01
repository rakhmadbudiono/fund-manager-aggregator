package fundmanager

type OtherFundManager struct {
	FundManager
}

func newOtherFundManager(fm FundManager) IFundManager {
	return &OtherFundManager{
		FundManager: fm,
	}
}
