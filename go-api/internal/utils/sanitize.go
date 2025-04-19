package utils

func MaskCardNumber(card string) string {
	if len(card) <= 4 {
		return card
	}
	return "************" + card[len(card)-4:]
}
