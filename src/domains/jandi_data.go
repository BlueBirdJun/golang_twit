package domains

type JandiData struct {
	Body         string        `json:"body"`
	ConnectColor string        `json:"connectColor"`
	ConnectInfo  []ConnectInfo `json:"connectInfo"`
}

type ConnectInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"imageUrl,omitempty"`
}
