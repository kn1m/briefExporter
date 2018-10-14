package exporters

type (
	NoteRecord struct {
		BookTitle        string   `json:"book_title"`
		BookOriginalName string   `json:"book_original_name"`
		CreatedOn        string   `json:"created_on"`
		NoteTitle        string   `json:"note_title"`
		NoteText         string   `json:"note_text"`
		BookAuthor       []Author `json:"authors"`
		FirstPage        int      `json:"first_page"`
		SecondPage       int      `json:"second_page"`
		FirstLocation    int      `json:"first_location"`
		SecondLocation   int      `json:"second_location"`
	}

	Author struct {
		FirstName     string `json:"first_name"`
		SecondaryName string `json:"secondary_name"`
		Surname       string `json:"surname"`
	}
)

type Exporter interface {
	GetNotes(path string) ([]*NoteRecord, error)
}