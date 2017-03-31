package main

// IViewModel Interface for ViewModels
type IViewModel interface {
	ToJson() string
	FromModel()
}
