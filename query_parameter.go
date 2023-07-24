package baseapp

import "strings"

type QueryParameter struct {
	Page   int
	Limit  int
	Search string
}

func (i QueryParameter) UpperCase() *QueryParameter {
	return &QueryParameter{
		Page:   i.Page,
		Limit:  i.Limit,
		Search: strings.ToUpper(i.Search),
	}
}
