package service

type NodeStore struct {
	kv map[string]Node
}

var (
	nodeStore = NewNodeStore()
)

func NewNodeStore() *NodeStore {
	return &NodeStore{
		kv: make(map[string]Node),
	}
}


// NodeStore
func (ns *NodeStore) addNode(n Node) {
	ns.kv[n.hash()] = n
}

func (ns *NodeStore) getNode(hash string) Node {
	return ns.kv[hash]
}

