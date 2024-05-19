package response

type Common struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    Meta        `json:"_meta,omitempty"`
}

type Meta struct {
	Page       string `json:"page,omitempty"`
	PerPage    string `json:"per_page,omitempty"`
	PageCount  string `json:"page_count,omitempty"`
	TotalCount string `json:"total_count,omitempty"`
	Links      Links  `json:"links,omitempty"`
}
type Links struct {
	Self     string `json:"self,omitempty"`
	First    string `json:"first,omitempty"`
	Previous string `json:"previous,omitempty"`
	Next     string `json:"next,omitempty"`
	Last     string `json:"last,omitempty"`
}
