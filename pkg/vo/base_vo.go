package vo

import "encoding/json"

type IDNamePair struct {
	Id   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func IDNamePairOf(id int64, name string) *IDNamePair {
	return &IDNamePair{
		Id:   id,
		Name: name,
	}
}

func (p *IDNamePair) String() string {
	if j, err := json.Marshal(p); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

type IDNameEmailTriple struct {
	Id    int64  `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func IDNameEmailTripleOf(id int64, name string, email string) *IDNameEmailTriple {
	return &IDNameEmailTriple{
		Id:    id,
		Name:  name,
		Email: email,
	}
}

func (t *IDNameEmailTriple) String() string {
	if j, err := json.Marshal(t); err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
