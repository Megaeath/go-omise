package processor

import (
	"errors"
	"fmt"
	"mockdonate/internal/model"
)

func MockCharge(d model.Donation) error {
	// Clear credit card data after use
	defer func() {
		d.CCNumber = ""
		d.CVV = ""
	}()

	if d.CVV == "999" {
		return errors.New("charge declined due to invalid CVV")
	}

	fmt.Printf("✅ Charged %s: %.2f THB\n", d.Name, float64(d.AmountSubunits)/100)
	return nil
}
