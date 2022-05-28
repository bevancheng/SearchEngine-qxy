package model

type SearchRequest struct {
	Query string `json:"query,omitempty"`
	Order string `json:"order,omitempty"`
	Page  int    `json:"page,omitempty"`
	Limit int    `json:"limit,omitempty"`
}

func (s *SearchRequest) GetAndSetDefault() *SearchRequest {
	if s.Limit == 0 {
		s.Limit = 100
	}
	if s.Page == 0 {
		s.Page = 1
	}

	if s.Order == "" {
		s.Order = "desc"
	}
	return s
}

type SearchResult struct {
	Time      float64       `json:"time,omitempty"`
	Total     int           `json:"total"`
	PageCount int           `json:"pageCount"`
	Page      int           `json:"page,omitempty"`
	Limit     int           `json:"limit,omitempty"`
	Document  []ResponseDoc `json:"documents,omitempty"`
	Words     []string      `json:"words,omitempty"`
}
