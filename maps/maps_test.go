package maps

import "testing"

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		dictionary.Add(key, value)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})

}

func assertDefinition(t testing.TB, dictionary Dictionary, key, value string) {
	t.Helper()
	got, err := dictionary.Search("test")
	if err != nil {
		t.Fatal("should find added word: ", err)
	}
	if value != got {
		t.Errorf("got %q, want %q", got, value)
	}
}

func TestSearch(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)

	})
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "new definition"
		err := dictionary.Update(word, newDefinition)
		assertError(t, err, ErrWordDoesNotExist)

	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "this is just a test"}
	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("expected %q to be deleted", word)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
