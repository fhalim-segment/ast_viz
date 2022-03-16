package astviz

import (
	"io"
	"os"
	"testing"
)

func TestSimple(t *testing.T) {
	testFile("../examples/example1.json", t)
}
func TestHung1(t *testing.T) {
	testFile("../examples/example2.json", t)
}

func TestOptions(t *testing.T) {
	testFile("../examples/example3.json", t)
}

func testFile(fileName string, t *testing.T) {
	file, err := os.Open(fileName)
	defer file.Close()
	handleError(err, t)
	b, err := io.ReadAll(file)
	handleError(err, t)
	node, err := Load(b)
	handleError(err, t)
	if node == nil {
		t.Log("Got a null node")
		t.FailNow()
	}
	node.Dump(os.Stderr, 0)
}

func handleError(e error, t *testing.T) {
	if e != nil {
		t.Log(e)
		t.FailNow()
	}
}
