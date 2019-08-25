package main

// PostmanInfo is postman first item for info
type PostmanInfo struct {
	PostmanID   string `json:"_postman_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Schema      string `json:"schema"`
}

// Item struct which contains the Item
type Item struct {
	Info        *PostmanInfo `json:"info"`
	Name        string       `json:"name"`
	Request     *Request     `json:"request"`
	Description string       `json:"description"`
	Response    []*Response  `json:"response"`
	Event       Event        `json:"event"`
	Items       []*Item      `json:"item"`
	path        string
	parent      *Item
	level       int
}

// Request struct which contains the Item
type Request struct {
	Method      string    `json:"method"`
	Header      []*Header `json:"header"`
	Body        Body      `json:"body"`
	URL         URL       `json:"url"`
	Description string    `json:"description"`
	file        string
}

// Response struct which contains the Item
type Response struct {
	Name            string    `json:"name"`
	OriginalRequest Request   `json:"originalRequest"`
	Status          string    `json:"status"`
	Code            int       `json:"code"`
	Method          string    `json:"method"`
	Header          []*Header `json:"header"`
	Body            Body      `json:"body"`
	URL             URL       `json:"url"`
	Description     string    `json:"description"`
}

// Header struct
type Header struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	Value       string `json:"value"`
	Description string `json:"description"`
}

// Body struct
type Body struct {
	Mode        string `json:"mode"`
	Raw         string `json:"raw"`
	Description string `json:"description"`
}

// URL struct
type URL struct {
	Raw         string `json:"raw"`
	Host        string `json:"host"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

// Event struct
type Event struct {
	Listen string `json:"test"`
	Script Script `json:"script"`
}

// Script struct
type Script struct {
	ID   string   `json:"id"`
	Exec []string `json:"exec"`
	Type string   `json:"type"`
}

// Cookie struct
type Cookie struct {
	Expires  string `json:"expires"`
	HTTPOnly bool   `json:"httpOnly"`
	Domain   string `json:"domain"`
	Path     string `json:"path"`
	Secure   string `json:"secure"`
	Key      string `json:"key"`
	Value    string `json:"value"`
}
