package pointers

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	

	t.Run("test deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		
	})

	t.Run("test withdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err)
	})
}