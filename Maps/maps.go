package main

import "errors"




type Dictionary map[string]string

var ( ErrNotFound = errors.New("Could not find the word you are looking for")
errWordExist = errors.New("cannot add word because it already exists")
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