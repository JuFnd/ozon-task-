// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	Author    *User  `json:"author"`
	Post      *Post  `json:"post"`
	ParentID  string `json:"parent_id"`
	CreatedAt string `json:"created_at"`
}

type Mutation struct {
}

type Post struct {
	ID          string `json:"id"`
	Content     string `json:"content"`
	CreatedAt   string `json:"created_at"`
	Author      *User  `json:"author,omitempty"`
	IsCommented *bool  `json:"isCommented,omitempty"`
}

type Query struct {
}

type User struct {
	ID    string `json:"id"`
	Login string `json:"login"`
}