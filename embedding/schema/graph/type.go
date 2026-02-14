package graph

type ParentNode struct {
	NodeID    string
	ArticleID string
	ChunkID   string
	Title     string
	Tag       string
	Keywords  []string
}

type ChildNode struct {
	NodeID   string
	ChunkID  string
	Title    string
	Tag      string
	Keywords []string
}

type Edge struct {
	EdgeID     string
	FromNodeID string
	ToNodeID   string
	Weight     float64
	Tag        string
}
