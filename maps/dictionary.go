package main

// Dictionary is a type map
type Dictionary map[string]string

// DictionaryErr is a string holding an error
type DictionaryErr string

// Err* are error types for the dictionary funtions
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it doesn't exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search looks for a word in the dictionary and returns the definition
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add adds a value to the dictionary returns and error
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// Update changes the definiton of a word in the dictionary
func (d Dictionary) Update(word, definition string) error {
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

// Delete deletes a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
