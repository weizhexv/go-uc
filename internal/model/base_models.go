package model

type IdCountPair struct {
	id    int64
	count int64
}

func (p *IdCountPair) Id() int64 {
	return p.id
}

func (p *IdCountPair) Count() int64 {
	return p.count
}
