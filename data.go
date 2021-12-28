package main

type todo struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allTodos []todo

var todos = allTodos{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		ID:          "2",
		Title:       "Advanced Golang",
		Description: "Learn advanced concepts of Go",
	},
}
