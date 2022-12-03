package internal

import "testing"

func TestReader_ReadString(t *testing.T) {
	reader := NewStringReader("hello\nworld")
	if reader.Scan() != true {
		t.Fatalf("reader.Scan() != true")
	}
	if s := reader.ParseString(); s != "hello" {
		t.Fatalf("%v != hello", s)
	}
	if reader.Scan() != true {
		t.Fatalf("reader.Scan() != true")
	}
	if s := reader.ParseString(); s != "world" {
		t.Fatalf("%v != world", s)
	}
	if reader.Scan() != false {
		t.Fatalf("reader.Scan() != false")
	}
}

func TestReader_ReadString_WithDelimiter(t *testing.T) {
	reader := NewStringReader("hello\n\nworld").SetDelimiter("\n\n")
	if reader.Scan() != true {
		t.Fatalf("reader.Scan() != true")
	}
	if s := reader.ParseString(); s != "hello" {
		t.Fatalf("%v != hello", s)
	}
	if reader.Scan() != true {
		t.Fatalf("reader.Scan() != true")
	}
	if s := reader.ParseString(); s != "world" {
		t.Fatalf("%v != world", s)
	}
	if reader.Scan() != false {
		t.Fatalf("reader.Scan() != false")
	}
}

func TestReader_ReadInt(t *testing.T) {
	reader := NewStringReader("+1\n-12")
	if reader.Scan() != true {
		t.Fatalf("reader.Scan() != true")
	}
	if s := reader.ParseInt(); s != 1 {
		t.Fatalf("%v != 1", s)
	}
	if reader.Scan() != true {
		t.Fatalf("reader.Scan() != true")
	}
	if s := reader.ParseInt(); s != -12 {
		t.Fatalf("%v != -12", s)
	}
	if reader.Scan() != false {
		t.Fatalf("reader.Scan() != false")
	}
}
