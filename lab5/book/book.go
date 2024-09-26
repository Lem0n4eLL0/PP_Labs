package book

import (
	"PP_LABS/utils/stringutils"
)

type Book struct {
	title   string
	authors []string
	pages   []string
}

func NewBook(title string, authors []string, pages []string) *Book {
	return &Book{title: title, authors: authors, pages: pages}
}

func (b *Book) ToString() string {
	return "title:\n" + b.title +
		"\nautors:\n" + stringutils.StringBuilder(b.authors) +
		"pages:\n" + stringutils.StringBuilder(b.pages)
}
