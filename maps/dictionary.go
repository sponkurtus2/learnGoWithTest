package maps

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

/* Initizalize an empty directory like this
* var dict := make(map[string]string)
 */

// Global var
const (
	// ErrWordNotFound = errors.New("word doesnt't exist")
	// ErrWordExists   = errors.New("word already exist")
	ErrWordNotFound     = DictionaryErr("word not found")
	ErrWordExists       = DictionaryErr("word already exist")
	ErrWordDoesNotExist = DictionaryErr("Word does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrWordNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, content string) error {
	_, err := d.Search(key)

	switch err {
	case ErrWordNotFound:
		d[key] = content
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrWordNotFound:
		return ErrWordDoesNotExist
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
