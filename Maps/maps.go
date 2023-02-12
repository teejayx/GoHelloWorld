package main

//import "errors"




type Dictionary map[string]string
type DictionaryErr string 

const ( ErrNotFound = DictionaryErr("Could not find the word you are looking for")
errWordExist = DictionaryErr("cannot add word because it already exists")
errWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)




func (d Dictionary) Search(word string) (string, error) {
   definition, ok := d[word] 
   if !ok {
	return "", ErrNotFound
   }

   return definition, nil
}


func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil: 
	    return errWordExist
	default: 
	    return err
	}
	
	return nil
}


func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
   
	switch err {
	case ErrNotFound:
		return errWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	d[word] = definition
	return nil
}

func (e DictionaryErr) Error() string {
	return string(e)
}


func (d Dictionary) Delete(word string) {
  delete(d, word)
}