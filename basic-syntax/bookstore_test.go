package bookstore_test

import (
	"bookstore"
	"testing"
)

// c := bookstore.Customer {
// 	Name: "John Doe",
// 	Email: "a@example",
// }

// book := bookstore.Book {
// 	"hello world",
// 	"ajiea.chow",
// 	10
// }

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{ // struct 测试方法
		"hello world",
		"ajiea.chow",
		10,
	}
}
