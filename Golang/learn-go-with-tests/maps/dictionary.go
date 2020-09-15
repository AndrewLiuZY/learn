package main

const (
	ErrNotFound         = DictionatyErr("could not find the word you are looking for")
	ErrWordExists       = DictionatyErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionatyErr("connot update word because it did not exist")
)

type DictionatyErr string

func (e DictionatyErr) Error() string {
	return string(e)
}

type Dictionaty map[string]string

func (d Dictionaty) Search(key string) (string, error) {
	word, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return word, nil
}

func (d Dictionaty) Add(key, word string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = word
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionaty) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionaty) Delete(word string) {
	delete(d, word)
}
