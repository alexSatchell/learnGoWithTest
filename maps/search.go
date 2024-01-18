package search

const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// maps can return two values, the second value is a bool,
	// which indicates a successful find
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

/*
* Notice that there is not pointer to Dictionary,
* In GO, maps are LIKE reference types, BUT THEY ARE NOT
* when you pass a map to a function, you are indeed passing
* a copy of the map, but this copy includes the pointer to the underlying data structure.
 */
func (d Dictionary) Add(word string, definition string) error {
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
