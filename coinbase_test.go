package coinbase

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"

	response, _ := Get(fmt.Sprintf("%s-%s", wantCoin, wantCurrency))
	if response.Base != wantCoin {
		t.Errorf("Translate() = %q, want %q", response.Base, wantCoin)
	}

	if response.Currency != wantCurrency {
		t.Errorf("TestGet() = %q, want %q", response.Currency, wantCoin)
	}

}

func TestGetWithDate(t *testing.T) {

	wantCoin := "SHIB"
	wantCurrency := "EUR"
	wantAmount := 0.0000190

	response, _ := GetWithDate(fmt.Sprintf("%s-%s", wantCoin, wantCurrency), "2022-02-01")
	if response.Base != wantCoin {
		t.Errorf("TestGetWithDate() = %q, want %q", response.Base, wantCoin)
	}

	if response.Currency != wantCurrency {
		t.Errorf("TestGetWithDate() = %q, want %q", response.Currency, wantCoin)
	}

	if response.Amount != wantAmount {
		t.Errorf("TestGetWithDate() = %.7f, want %.7f", response.Amount, wantAmount)
	}

}
