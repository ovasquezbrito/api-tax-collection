package models

import "strings"

type QueryParameter struct {
	Page   int    `json:"page"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

func (i QueryParameter) UpperCase() *QueryParameter {
	return &QueryParameter{
		Page:   i.Page,
		Limit:  i.Limit,
		Search: strings.ToUpper(i.Search),
	}
}
