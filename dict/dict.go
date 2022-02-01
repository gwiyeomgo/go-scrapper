package dict

import "errors"

//Dictionary 는 alias(별명) 같은 거다
//Dictionary 는 map[string]string의 가명이다

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExists = errors.New("That word aleredy exists")
var errCantUpdate = errors.New("Cant update non-existing word")

func (d Dictionary) Search(word string) (string, error) {

	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a wrod
func (d Dictionary) Delete(word string) {
	//map 관련 글을 보면
	//delete함수를 사용해서 map에서 삭제 할 수 있다.
	delete(d, word)
}
