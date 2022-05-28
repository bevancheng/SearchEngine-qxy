package model

type IndexDoc struct {
	Id       uint64         `json:"id,omitempty"`
	Text     string         `json:"text,omitempty"`
	Document map[string]any `json:"document,omitempty"`
}

type StorageIndexDoc struct {
	*IndexDoc
	Keys []string `json:"keys,omitempty"`
}

type ResponseDoc struct {
	IndexDoc
	OriginalText string   `json:"originalText,omitempty"`
	Score        int      `json:"score,omitempty"`
	Keys         []string `json:"keys,omitempty"`
}

type RemoveIndexModel struct {
	Id uint64 `json:"id,omitempty"`
}
