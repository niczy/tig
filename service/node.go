package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
)


type Node struct {
	name string
	isDir bool
	// children nodes if isDir = true
	childrenHash []string
	// hash of the content if isDir = false
	contentHash string
}

type NodeUpdateType string

const (
	OpDelete NodeUpdateType = "delete"
	OpUpdate NodeUpdateType = "update"
	OpAdd NodeUpdateType = "add"
)




func NewRoot(name string) *Node {
	return &Node{
		name: name,
		isDir: true,
		childrenHash : make([]string, 0, 5),
	}
}

func PrintTree(node *Node) {
	PrintTreeInternal(node)
	fmt.Println()
}

func PrintTreeInternal(node *Node) {
	fmt.Printf("%s", node.name)
	if node.isDir {
		fmt.Printf("(")
		for _, c := range node.childrenHash {
			cNode := nodeStore.getNode(c)
			PrintTreeInternal(&cNode)
		}
		fmt.Printf(")")
	}
}


// Node
func (n *Node) hash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", *n)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (n *Node) listFiles(paths []string) ([]string, error) {
	leaf := n.getNodeByPaths(paths)
	if leaf == nil {
		return nil, errors.New("path not found")
	}
	if !leaf.isDir {
		return nil, errors.New("path is not a directory")
	}
	// Build list of files.
	ret := make([]string, 0, leaf.childrenLength())
	for i:=0; i < leaf.childrenLength(); i++ {
		c := leaf.child(i)
		ret = append(ret, c.name)
	}
	return ret, nil
}

func (n *Node) childrenLength() int {
	return len(n.childrenHash)
}

func (n *Node) child(idx int) *Node {
	c := nodeStore.getNode(n.childrenHash[idx])
	return &c
}

func (n *Node) childByName(name string) *Node {
	for _, ch := range n.childrenHash {
		c := nodeStore.getNode(ch)
		if c.name == name {
			return &c
		}
	}
	return nil
}

func (n *Node) getNodeByPaths(paths []string) *Node {
	for _, path := range paths {
		if path == "" {
			continue
		}
		exist := false
		for i:=0; i < n.childrenLength(); i++ {
			if path == n.child(i).name {
				n = n.child(i)
				exist = true
				break
			}
		}
		if !exist {
			return nil
		}
	}
	return n
}

// Apply UpdateNode tree to current tree.
func (n *Node) applyUpdateNodes(un *UpdateNode) {
	if n.hash() != un.hash && un.updateType != OpAdd {
		log.Fatalf("Applying update with hash %s to node with hash %s", un.hash, n.hash())
	}
	childHash := make([]string, 0, len(n.childrenHash))
	if n.isDir {
		// Update directory node.
		for _, ch := range n.childrenHash {
			if _, ok := un.childUpdateNodes[ch]; ok {
				u := un.childUpdateNodes[ch]
				if u.updateType == OpUpdate {
					c := nodeStore.getNode(ch)
					c.applyUpdateNodes(u)
					childHash = append(childHash, c.hash())
				}
			} else {
				// No update needed, append original child.
				childHash = append(childHash, ch)
			}
		}
	} else {
		// Update file node.
		n.name = un.name
		n.isDir = un.isDir
		n.contentHash = un.contentHash
	}
	// Process Add operation
	for _, cun := range un.childUpdateNodes {
		if cun.updateType == OpAdd {
			c := &Node{
				name: cun.name,
				isDir: cun.isDir,
				contentHash: cun.contentHash,
			}
			if cun.isDir {
				c.applyUpdateNodes(cun)
			}
			nodeStore.addNode(*c)
			log.Printf("New child node hash = %s", c.hash())
			childHash = append(childHash, c.hash())
		}
	}
	n.childrenHash = childHash
	log.Printf("childrenHash %v", n.childrenHash)
	nodeStore.addNode(*n)
}

