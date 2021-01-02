package service

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"log"
	"os"
	"strconv"
)

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestListFiles(t *testing.T) {
	root := "./test_data"
	os.RemoveAll(root)
	os.Mkdir(root, os.ModePerm)

	defer os.RemoveAll(root)

	tig := NewTig()
	files, err := tig.ListFiles("/")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, []string{}, files)

	for i:=0; i < 10; i++ {
		tig.UpdateFile(strconv.Itoa(i)+".txt", "content")
	}

	expectedFiles := []string{"0.txt", "1.txt", "2.txt", "3.txt", "4.txt", "5.txt", "6.txt", "7.txt", "8.txt", "9.txt"}

	files, _ = tig.ListFiles("/")
	assert.Equal(t, expectedFiles, files)

	str := "Hello world!"
	tig.UpdateFile("h.txt", str)

	tmpData, _ := tig.ReadFile("h.txt")
	assert.Equal(t, tmpData, str)
}

func TestDiff(t *testing.T) {
	text1 := "hello world"
	text2 := "hellp baba"
	dmp := diffmatchpatch.New()
	diffs := dmp.DiffMain(text1, text2, false)
	fmt.Println(dmp.DiffPrettyText(diffs))
}
