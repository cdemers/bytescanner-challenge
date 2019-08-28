package bytescanner

import (
	"bytes"
	"fmt"
	"io"
	"testing"
)

var (
	scanner *ByteScanner
	reader  io.Reader
)

func ExampleByteScanner_Eat() {
	bs := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	reader = bytes.NewReader(bs)
	scanner = NewByteScanner(reader)
	//scanner.Open()

	b, err := scanner.Eat()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%d", b)
	// Output: 0
}

func TestPeek0AndThanEat(t *testing.T) {
	reader = bytes.NewReader([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
	scanner = NewByteScanner(reader)

	b, err := scanner.Peek(0)
	if err != nil {
		t.Error("scanner.Peek(0) expected no error, got:", err)
	}
	if b != 0 {
		t.Errorf("scanner.Peek(0) expected to return 0, got %d", b)
	}

	b, err = scanner.Eat()
	if err != nil {
		t.Error("scanner.Eat() expected no error, got:", err)
	}
	if b != 0 {
		t.Errorf("scanner.Eat() expected to return 0, got %d", b)
	}
}

func TestEatUntilEOF(t *testing.T) {
	bs := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	reader = bytes.NewReader(bs)
	scanner = NewByteScanner(reader)

	for ctr := 0; ctr < len(bs); ctr++ {
		b, err := scanner.Eat()
		if err != nil {
			t.Errorf("[CTR:%d B:%d] expected no error, got: %s", ctr, b, err)
		}
		if b != byte(ctr) {
			t.Errorf("expected Eat() to return %d, got %d", ctr, b)
		}
	}

	_, err := scanner.Eat()
	if err == nil {
		t.Error("expected io.EOF error, got no error")
	}

}

func TestPeekUntilEOF(t *testing.T) {
	bs := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	reader = bytes.NewReader(bs)
	scanner = NewByteScanner(reader)
	// scanner.Open()

	maximumPeekOffset := 8 // This should be configured according to your implementation.
	var ctr int
	for ctr = 0; ctr < maximumPeekOffset; ctr++ {
		b, err := scanner.Peek(ctr)
		if err != nil {
			t.Error("expected no error, got:", err)
		}
		if b != byte(ctr) {
			t.Errorf("expected Peak(%d) to return %d, got %d", ctr, ctr, b)
		}
	}

	_, err := scanner.Peek(ctr + 1)
	if err == nil {
		t.Error("expected io.EOF error, got no error")
	}

}

func TestPeek0AndEatUntilEOF(t *testing.T) {
	bs := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	reader = bytes.NewReader(bs)
	scanner = NewByteScanner(reader)

	var ctr int
	for ctr = 0; ctr < len(bs); ctr++ {

		b, err := scanner.Peek(0)
		if err != nil {
			t.Errorf("scanner.Peek(0) (ctr=%d) expected no error, got: %s", ctr, err)
		}
		if b != byte(ctr) {
			t.Errorf("scanner.Peek(0) expected to return %d, got %d", ctr, b)
		}

		b, err = scanner.Eat()
		if err != nil {
			t.Errorf("scanner.Eat() (ctr=%d) expected no error, got: %s", ctr, err)
		}
		if b != byte(ctr) {
			t.Errorf("scanner.Eat() expected to return %d, got %d", ctr, b)
		}
	}

	b, err := scanner.Peek(ctr + 1)
	if err == nil {
		t.Errorf("scanner.Peek(ctr + 1) (ctr+1=%d) (b=%d) expected io.EOF error, got error no error", ctr+1, b)
	}

	_, err = scanner.Eat()
	if err != io.EOF {
		t.Error("scanner.Eat() expected io.EOF error, got no error")
	}

}
