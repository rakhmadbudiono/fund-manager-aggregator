package mutualfund

type OtherMutualFund struct {
	MutualFund
}

func newOtherMutualFund(mf MutualFund) IMutualFund {
	return &OtherMutualFund{
		MutualFund: mf,
	}
}
