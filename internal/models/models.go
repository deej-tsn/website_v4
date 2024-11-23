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
		Name   string
		ImgURL string
	}

	ContactIcon struct {
		ImgURL string
		URL    string
		Name   string
	}

	ProjectHero struct {
		Title       string
		Description string
		ImgURL      string
		Skills      [][]string
		Type        string
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

func NewSkill(name string, url string) *SkillIcon {
	skill := SkillIcon{
		Name:   name,
		ImgURL: fmt.Sprintf("fa-brands fa-%s fa-2xl", url),
	}
	return &skill
}
