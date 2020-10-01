package cyoa

// Declare structs to host JSON elements
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// Declare map type to host JSON data
type Story map[string]Chapter
