package model

type ChargeRequest struct {
	Name           string `json:"name" binding:"required"`
	AmountSubunits int    `json:"amount_subunits" binding:"required,gt=0"`
	CCNumber       string `json:"cc_number" binding:"required,len=16"`
	CVV            string `json:"cvv" binding:"required,len=3"`
	ExpMonth string `json:"exp_month" binding:"required,min=1,max=2"`
	ExpYear        string `json:"exp_year" binding:"required,len=4"`
}

type ChargeResponse struct {
	LogID string `json:"log_id"`
}
