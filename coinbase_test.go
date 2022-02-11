package coinbase

import (
	"fmt"
	"testing"
)

func TestTranslate(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"

	response, _ := Get(fmt.Sprintf("%s-%s", wantCoin, wantCurrency))

	fmt.Printf("%.7f\n", response.Amount)

	if response.Base != wantCoin {
		t.Errorf("Translate() = %q, want %q", response.Base, wantCoin)
	}

	if response.Currency != wantCurrency {
		t.Errorf("Translate() = %q, want %q", response.Currency, wantCoin)
	}

}
