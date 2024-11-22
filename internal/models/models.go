package models

import (
	"fmt"
	"time"
)

type (
	BlogResp struct {
		ID    string `query:"id"`
		Title string `query:"title"`
		Body  string `query:"body"`
	}

	Post struct {
		Title       string    `toml:"title"`
		Slug        string    `toml:"slug"`
		Description string    `toml:"description"`
		Date        time.Time `toml:"date"`
		Author      Author    `toml:"author"`
	}

	Author struct {
		Name  string `toml:"name"`
		Email string `toml:"email"`
	}

	SkillIcon struct {
		Name       string
		URL        string
		Experience string
	}

	ContactIcon struct {
		ImgURL string
		URL    string
		Name   string
	}
)

func NewContact(name string, pathToLink string, url string) *ContactIcon {
	contact := ContactIcon{
		Name:   name,
		ImgURL: fmt.Sprintf("%s/%s.svg", pathToLink, name),
		URL:    url,
	}
	return &contact
}
