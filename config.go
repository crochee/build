package main

type Config struct {
	Product *Product `json:"product" yaml:"product"`
	Script  string   `json:"script" yaml:"script"`
}

type Product struct {
	Type string
	Name string
}
