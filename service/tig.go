package service

import (
	"io/ioutil"
	"path/filepath"
	"log"
)

type Tig struct {
	root string
}

func NewTig(root string) (t *Tig) {
  tig := &Tig{
  	root: root,
  }
  return tig
}

func (t *Tig) ListFiles(dir string) []string {
        files, err := ioutil.ReadDir(filepath.Join(t.root, dir))
	if err != nil {
		log.Fatal(err)
	}

	ret := make([]string, 0, 0)
        for _, f := range files {
		ret = append(ret, f.Name())
        }
	return ret 
}

func (t *Tig) ReadFile(file string) string {
        ret, _ := ioutil.ReadFile(filepath.Join(t.root, file))	
	return string(ret)
}

func (t *Tig) GetRoot() string {
	return t.root
}


