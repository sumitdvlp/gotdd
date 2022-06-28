package hello

import "testing"

func TestHello(t *testing.T) {
	//  In Go you can declare functions inside other functions and assign them to variables.
	// You can then call them, just like normal functions.
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function
		// call rather than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("Got %q want %q", got, want)
		}
	}

	t.Run("saying hello to actual people", func(t *testing.T) {
		got := Hello("Sumi", "")
		want := "Hello, Sumi"

		assertCorrectMessage(t, got, want)

	})

	t.Run("saying hello to empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Lahm", "French")
		want := "Bonjour, Lahm"
		assertCorrectMessage(t, got, want)
	})
}
