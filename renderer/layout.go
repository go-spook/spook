package renderer

import (
	"html/template"

	"github.com/go-spook/spook/model"
)

// Layout is the base layout of the website
type Layout struct {
	WebsiteTitle  string
	WebsiteOwner  string
	ContentTitle  string
	ContentDesc   string
	ContentAuthor string
	Pages         []model.Page
}

// List is layout that used in list template.
// Might be used in frontpage template as well.
type List struct {
	Layout
	Type        ListType
	Path        string
	Posts       []model.Post
	Tags        []model.Group
	Categories  []model.Group
	CurrentPage int
	MaxPage     int
}

// Page is layout that used in single page
type Page struct {
	Layout
	Thumbnail string
	HTML      template.HTML
}

// Post is layout that used in post
type Post struct {
	Layout
	CreatedAt string
	UpdatedAt string
	Category  model.Group
	Tags      []model.Group
	Thumbnail string
	HTML      template.HTML
	Older     model.Post
	Newer     model.Post
}
