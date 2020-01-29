package models

type Data struct {
	Name     string
	URL      string
	Language []string
	Lines    []int
}

type Filter struct {
	Filter string
}
