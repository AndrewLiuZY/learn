package main

import "testing"

func TestSearch(t *testing.T) {

	dic := Dictionaty(map[string]string{"test": "just a test"})

	t.Run("Known word", func(t *testing.T) {
		want := "just a test"

		assertDefinition(t, dic, "test", want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dic.Search("hello")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("Add new word", func(t *testing.T) {
		dic := Dictionaty{}
		dic.Add("test", "I am a test value")

		want := "I am a test value"

		assertDefinition(t, dic, "test", want)
	})

	t.Run("Add existing word", func(t *testing.T) {
		word := "test"
		definition := "just a test"

		dic := Dictionaty{word: definition}
		err := dic.Add(word, definition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dic, word, definition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update a exists word", func(t *testing.T) {
		word := "test"
		definition := "just test"
		dic := Dictionaty{word: definition}
		newDefinition := "just a new test"
		dic.Update(word, newDefinition)

		assertDefinition(t, dic, word, newDefinition)
	})

	t.Run("update a new word", func(t *testing.T) {
		dic := Dictionaty{}
		err := dic.Update("test", "new test")

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dictionaty{word: "just test"}

	dic.Delete(word)
	_, err := dic.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected '%s' to be deleted", word)
	}

}

func assertDefinition(t *testing.T, dic Dictionaty, word, definition string) {
	t.Helper()
	got, err := dic.Search(word)

	if err != nil {
		t.Fatal("should find added word")
	}
	if got != definition {
		t.Errorf("got %s want %s", got, definition)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}
}
