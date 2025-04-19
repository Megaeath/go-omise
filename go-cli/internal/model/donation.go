package model

type DonationRow struct {
	Name           string `json:"name" `
	AmountSubunits int    `json:"amount_subunits"`
	CCNumber       string `json:"cc_number" `
	CVV            string `json:"cvv"`
	ExpMonth       string `json:"exp_month" `
	ExpYear        string `json:"exp_year" `
}

type DonationResult struct {
	Row     DonationRow
	Success bool
	Error   error
}
