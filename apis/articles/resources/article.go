package resources

type Article struct {
	Id      string `json:"id,omitempty"`
	Author  string `json:"author,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}
