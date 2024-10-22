// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        string   `json:"id" gorm:"primary_key"`
	PostID    string   `json:"postID"`
	ParentID  *string  `json:"parentID,omitempty"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type Mutation struct {
}

type NewComment struct {
	PostID   string  `json:"postID"`
	ParentID *string `json:"parentID,omitempty"`
	Content  string  `json:"content"`
}

type NewPost struct {
	Title           string `json:"title"`
	Content         string `json:"content"`
	CommentsAllowed bool   `json:"commentsAllowed"`
}

type Post struct {
	ID              string       `gorm:"primaryKey" json:"id"`
	Title           string     `json:"title"`
	Content         string     `json:"content"`
	CommentsAllowed bool       `json:"commentsAllowed"`
	Comments        []*Comment `json:"comments" gorm:"foreignkey:PostID"`
	CreatedAt       string     `json:"createdAt"`
}

type Query struct {
}
