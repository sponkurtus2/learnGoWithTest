package pointerserrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertMsg := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()

		if got != want {
			t.Errorf("Got -> %s, Want -> %s", got, want)
		}
	}

	assertErr := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("Wanted an error, but didn't get one")
		}

		if got.Error() != want {
			t.Errorf("Got -> %s, Want -> %s", got, want)
		}
	}

	t.Run("Deposit method", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		assertMsg(t, wallet, Bitcoin(10))
	})

	t.Run("Whitdraw money", func(t *testing.T) {
		wallet := Wallet{money: 100}
		wallet.Whithdraw(50)

		want := Bitcoin(50)

		assertMsg(t, wallet, want)
	})

	t.Run("Whithdraw with insufficient amount", func(t *testing.T) {
		startingMoney := Bitcoin(10)

		wallet := Wallet{money: startingMoney}
		err := wallet.Whithdraw(Bitcoin(100))

		assertErr(t, err, "Cant Whitdraw, insufficient ammount")
		assertMsg(t, wallet, startingMoney)
	})
}
