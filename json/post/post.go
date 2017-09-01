package post

//Post defines the main post for a blog or similar
type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

//Author defines an author for a Post
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//Comment defines a comment for a Post
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
