package utils

import (
	"math"
)

type Paging struct {
	Page       int
	Limit      int
	Offset     int
	TotalPages int
	Count      int
}

var DefaultLimt = 10

func NewPaging(count, limit int) *Paging {
	return &Paging{Count: count, Limit: limit}
}

func (p *Paging) SetPage(page int) *Paging {
	p.Page = page
	return p
}

func (p *Paging) Calc() *Paging {
	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit < 1 {

		p.Limit = DefaultLimt
	}

	if p.Count < p.Limit {
		p.TotalPages = 1
	} else {
		p.TotalPages = int(math.Ceil(float64(p.Count) / float64(p.Limit)))
	}

	if p.Page > p.TotalPages {
		p.Page = p.TotalPages
	}

	p.Offset = (p.Page - 1) * p.Limit

	return p
}
