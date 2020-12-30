package service

import (
	"os"
	"strconv"
	"path/filepath"
	"io/ioutil"
)


import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewTig(t *testing.T) {
	tig := NewTig("./test_data")
	assert.Equal(t, tig.GetRoot(), "./test_data")
}

func TestListFiles(t *testing.T) {
	root := "./test_data"
	os.RemoveAll(root)
	os.Mkdir(root, os.ModePerm)

	defer os.RemoveAll(root)

	tig := NewTig("./test_data")
	files := tig.ListFiles(".")
	assert.Equal(t, files, []string{})

	for i:=0; i < 10; i++ {
		os.Create(filepath.Join(root, strconv.Itoa(i)+".txt"))
	}

	expectedFiles := []string{"0.txt", "1.txt", "2.txt", "3.txt", "4.txt", "5.txt", "6.txt", "7.txt", "8.txt", "9.txt"}

	files = tig.ListFiles(".")
	assert.Equal(t, files, expectedFiles)

	str := "Hello world!"
	ioutil.WriteFile(filepath.Join(root, "h.txt"), []byte(str), 0644)

	tmpData := tig.ReadFile("h.txt")
	assert.Equal(t, tmpData, str)
}
