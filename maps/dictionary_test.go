package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"Carlos": "Fernanda"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("Carlos")
		want := "Fernanda"

		assertString(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, got := dictionary.Search("Carlitos feo")

		assertErr(t, got, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		word := "Carlos"
		definition := "Fernanda"

		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "Fernanda dos")

		assertErr(t, err, ErrWordNotFound)
		asserDefinition(t, dictionary, word, definition)
	})

	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "Linux"
		definition := "Carlitos"

		err := dictionary.Add(key, definition)

		assertErr(t, err, nil)
		asserDefinition(t, dictionary, key, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertErr(t, err, nil)
		asserDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErr(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrWordNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func asserDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word: ", err)
	}

	assertString(t, got, definition)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}

func assertErr(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
