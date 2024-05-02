package news

type Filters struct {
	Type int `query:"type"`
}

type Detail struct {
	ID   string `param:"id"`
}