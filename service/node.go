package service

import (
	"crypto/sha256"
	"fmt"
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
	Delete NodeUpdateType = "delete"
	Update NodeUpdateType = "update"
	Add NodeUpdateType = "add"
)

type UpdateNode struct {
	hash string
	updateType NodeUpdateType
	isDir bool
	name string
	// children nodes to be updated if isDir = true
	childUpdateNodes map[string]*UpdateNode
	// hash of the content if updateType is Update or Add.
	contentHash string
}

type NodeStore struct {
	kv map[string]Node
}

var (
	nodeStore NodeStore
)

// NodeStore
func (ns *NodeStore) addNode(n Node) {
	ns.kv[n.hash()] = n
}

func (ns *NodeStore) getNode(hash string) Node {
	return ns.kv[hash]
}


// Node
func (n *Node) hash() string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", *n)))

	return fmt.Sprintf("%x", h.Sum(nil))
}

// Apply UpdateNode tree to current tree.
func (n *Node) applyUpdateNodes(un *UpdateNode) {
	childHash := make([]string, 0, len(n.childrenHash))
	for i, ch := range n.childrenHash {
		if _, ok := un.childUpdateNodes[ch]; ok {
			u := un.childUpdateNodes[ch]
			if u.updateType != Delete {
				c := nodeStore.getNode(ch)
				c.applyUpdateNodes(u)
				childHash = append(childHash, c.hash())
			}
		} else {
			childHash = append(childHash, ch)
		}
	}
	nodeStore.addNode(*n)
}

