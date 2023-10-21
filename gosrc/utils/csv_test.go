package utils

import "testing"

func TestCSVParse(t *testing.T) {
	type Person struct {
		Id   int    `csv:"Id"`
		Name string `csv:"Name"`
	}

	data := []byte("Id,Name\n1,John\n2,Paul")
	dest := []Person{}

	err := CSVParse(data, &dest)
	if err != nil {
		t.Fatalf("Error parsing CSV: %v", err)
	}

	expected := []Person{
		{Id: 1, Name: "John"},
		{Id: 2, Name: "Paul"},
	}

	if len(dest) != len(expected) {
		t.Fatalf("Expected %v, got %v", len(expected), len(dest))
	}

	for i := range dest {
		if dest[i] != expected[i] {
			t.Fatalf("Expected %v, got %v", expected[i], dest[i])
		}
	}
}
