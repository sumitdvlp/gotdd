package wallet

import (
	"testing"
)

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestWallet(t *testing.T) {

	assertCorrectMessage := func(t testing.TB, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("Got %v want %v", got, want)
		}
	}
	wallet := Wallet{}
	t.Run("deposit", func(t *testing.T) {

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)
		assertCorrectMessage(t, got, want)
	})

	t.Run("withdrawl", func(t *testing.T) {

		err := wallet.Withdrawl(Bitcoin(30))

		got := wallet.Balance()
		//fmt.Printf("address of balance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)
		assertCorrectMessage(t, got, want)

		assertError(t, err, ErrInsufficientFunds)

	})
}
