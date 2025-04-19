package model

type DonationRow struct {
	Name           string
	AmountSubunits int64
	CCNumber       string
	CVV            string
	ExpMonth       string
	ExpYear        string
}

type DonationResult struct {
	Row     DonationRow
	Success bool
	Error   error
}
