package models

type _Store struct {
}

var Store _Store

func Init() {
	Store = _Store{}
}

func (s *_Store) GetMessage() string {
	return "Hello World"
}
