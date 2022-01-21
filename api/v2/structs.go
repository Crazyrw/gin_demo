package v2

type File struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Len     int    `json:"len"`
	Content []byte `json:"-"`
}

type Files struct {
	Files []*File `json:"files"`
	Len   int     `json:"len"`
}

var id int
var files = Files{Files: []*File{}}
