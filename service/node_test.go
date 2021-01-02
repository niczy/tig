package service

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNodeUpdate(t *testing.T) {
	n := NewRoot("root")
	fmt.Println(n.hash())
	PrintTree(n)
	// Add file and a directory.
	addFile := &UpdateNode {
		hash: "1234",
		name: "1234",
		updateType: OpAdd,
		isDir: false,
		contentHash: "contentHash",
	}
	addDir := &UpdateNode {
		hash: "blabla",
		name: "directory",
		updateType: OpAdd,
		isDir: true,
		childUpdateNodes: map[string]*UpdateNode{
			"1234": addFile,
		},
	}
	un := &UpdateNode {
		hash: n.hash(),
		updateType: OpAdd,
		isDir: true,
		childUpdateNodes: map[string]*UpdateNode{
			"1234": addFile,
			"blabla": addDir,
		},
	}
	n.applyUpdateNodes(un)
	fmt.Printf("Updated hash %s\n", n.hash())
	PrintTree(n)
	rootHash := n.hash()
	// root(directory(1234)1234)
	assert.Equal(t, n.name, "root")
	assert.Equal(t, n.isDir, true)
	assert.Equal(t, 2, len(n.childrenHash))
	c1 := nodeStore.getNode(n.childrenHash[0])

	assert.Equal(t, "1234", c1.name)
	assert.Equal(t, false, c1.isDir)
	assert.Equal(t, "contentHash", c1.contentHash)

	c2 := nodeStore.getNode(n.childrenHash[1])

	assert.Equal(t, "directory", c2.name)
	assert.Equal(t, c2.isDir, true)
	assert.Equal(t, 1, len(c2.childrenHash))

	c3 := nodeStore.getNode(c2.childrenHash[0])
	assert.Equal(t, c3.isDir, false)
	assert.Equal(t, c3.name, "1234")

	// Update file.
	updateFile := &UpdateNode{
		hash: c1.hash(),
		updateType: OpUpdate,
		name: "4321",
		isDir: false,
		contentHash: "newContentHash",
	}
	un = &UpdateNode {
		hash: n.hash(),
		updateType: OpUpdate,
		isDir: true,
		childUpdateNodes: map[string]*UpdateNode{
			c1.hash(): updateFile,
		},
	}
	n.applyUpdateNodes(un)

	c1 = nodeStore.getNode(n.childrenHash[0])
	assert.Equal(t, "4321", c1.name)
	assert.Equal(t, false, c1.isDir)
	assert.Equal(t, "newContentHash", c1.contentHash)

	// Delete file.
	deleteFile := &UpdateNode{
		hash: c1.hash(),
		updateType: OpDelete,
	}
	deleteDir := &UpdateNode{
		hash: c2.hash(),
		updateType: OpDelete,
	}
	un = &UpdateNode {
		hash: n.hash(),
		updateType: OpUpdate,
		isDir: true,
		childUpdateNodes: map[string]*UpdateNode{
			c1.hash(): deleteFile,
			c2.hash(): deleteDir,
		},
	}
	n.applyUpdateNodes(un)
	PrintTree(n)

	assert.Equal(t, len(n.childrenHash), 0)

	// Tree before deletion still exists.
	rootV1 := nodeStore.getNode(rootHash)
	PrintTree(&rootV1)
}
