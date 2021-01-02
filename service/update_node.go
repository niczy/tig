package service

type UpdateNode struct {
	// Hash of the node to apply this update operation.
	hash string
	updateType NodeUpdateType
	isDir bool
	name string
	// children nodes to be updated if isDir = true
	childUpdateNodes map[string]*UpdateNode
	// hash of the content if updateType is Update or Add.
	contentHash string
}
