package model

type Donation struct {
	Name           string `csv:"Name"`
	AmountSubunits int    `csv:"AmountSubunits"`
	CCNumber       string `csv:"CCNumber"`
	CVV            string `csv:"CVV"`
	ExpMonth       int    `csv:"ExpMonth"`
	ExpYear        int    `csv:"ExpYear"`
}
