package common

type QueryParameter struct {
	Page             int
	Limit            int
	Offset           int
	SearchQueryParam *SearchQueryParam
}

type SearchQueryParam struct {
	Key   string
	Value string
}
