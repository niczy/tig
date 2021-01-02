package service

import (
	"errors"
	"log"
	"strings"
)


// Tig provides file interface backed by Node implementation.
type Tig struct {
	root *Node
}

const (
	Separator = "/"
)

func NewTig() (t *Tig) {
  tig := &Tig{
  	root: NewRoot("root"),
  }
  return tig
}

func (t *Tig) ListFiles(path string) ([]string, error) {
	paths := strings.Split(path, Separator)
	log.Printf("Paths are %v\n", paths)
	return t.root.listFiles(paths)
}

func (t *Tig) ReadFile(path string) (string, error) {
	paths := strings.Split(path, Separator)
	node := t.root.getNodeByPaths(paths)
	if node == nil {
		return "", errors.New("path doesn't exist")
	}
	return contentStore.Get(node.contentHash), nil

}

func (t *Tig) UpdateFile(path, content string) error {
	contentHash := contentStore.Put(content)
	paths := strings.Split(path, Separator)
	un := &UpdateNode {
		hash: t.root.hash(),
		name: t.root.name,
		isDir: t.root.isDir,
		updateType: OpUpdate,
		childUpdateNodes: make(map[string]*UpdateNode),
	}
	rootUn := un
	node := t.root
	for i, p := range paths {
		if node != nil {
			node = node.childByName(p)
		}
		var newUn *UpdateNode
		if node != nil {
			newUn = &UpdateNode {
				hash: node.hash(),
				name: node.name,
				isDir: node.isDir,
				updateType: OpUpdate,
			}
			un.childUpdateNodes[node.hash()] = newUn
		} else {
			newUn = &UpdateNode {
				hash: "",
				name: path,
				isDir: i < len(paths)-1,
				updateType: OpAdd,
			}
			un.childUpdateNodes[path] = newUn
		}

		un = newUn
	}
	un.contentHash = contentHash

	t.root.applyUpdateNodes(rootUn)
	return nil
}



