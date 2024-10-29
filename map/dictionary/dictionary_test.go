package practice

import (
	"testing"
)

func TestMap(t *testing.T) {
	dictionary := Dictionary{"key1": "value1"}

	t.Run("known word", func(t *testing.T) {
		word, _ := dictionary.Search("key1")
		want := "value1"

		assertStrings(t, word, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("invalid key")

		assertErrors(t, err, ErrWordNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, definition)

		assertErrors(t, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "key"
		definition := "value"
		updateDefinition := "update value"
		dictionary := Dictionary{word: word, definition: definition}

		err := dictionary.Update(word, updateDefinition)

		assertErrors(t, err, nil)
		assertDefinition(t, dictionary, word, updateDefinition)
	})

	t.Run("when word not existing", func(t *testing.T) {
		word := "key"
		definition := "value"

		dictionary := Dictionary{word: word, definition: definition}

		err := dictionary.Update("invalid", definition)

		assertErrors(t, err, ErrWordNotFound)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	assertErrors(t, err, ErrWordNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given", got, want)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Fatalf("got error %q want %q", got, want)
	}
}
