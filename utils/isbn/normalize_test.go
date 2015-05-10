package isbn

import (
	"testing"
)

func TestISBN13Check(t *testing.T) {
	// Unhypened version
	isbn := "9787115253057"
	if Normalize(isbn) != isbn {
		t.Errorf("Failed to normalize %s: valid expected, invalid found.", isbn)
	}
	// Hypened version
	isbn_hypened := "978-7-115-25305-7"
	if r := Normalize(isbn_hypened); r != isbn {
		t.Errorf("Failed to normalize %s: %s expected, %s found", isbn_hypened, isbn, r)
	}
	// Wrong validate digit
	wrong_valid := "9787115253059"
	if Normalize(wrong_valid) != "" {
		t.Errorf("Failed to normalize %s: invalid expected, valid found.", wrong_valid)
	}
}

func TestISBN10Check(t *testing.T) {
	// Unhypened version
	isbn := "7115253056"
	expected := "9787115253057"
	if r := Normalize(isbn); r != expected {
		t.Errorf("Failed to normalize %s: %s expected, %s found", isbn, expected, r)
	}
	// Hypened version
	isbn_hypened := "7-115-25305-6"
	if r := Normalize(isbn_hypened); r != expected {
		t.Errorf("Failed to normalize %s: %s expected, %s found", isbn_hypened, expected, r)
	}
	// Validate code X
	isbnX := "711525317X"
	expected = "9787115253170"
	if r := Normalize(isbnX); r != expected {
		t.Errorf("Failed to normalize %s: %s expected, %s found", isbnX, expected, r)
	}
	// Wrong validate digit
	wrong_valid := "711525305X"
	if Normalize(wrong_valid) != "" {
		t.Errorf("Failed to normalize %s: invalid expected, valid found.", wrong_valid)
	}
}
