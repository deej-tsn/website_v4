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
		ImgFill     string
		Skills      [][]string
		Link        string
		GithubLink  string
		Type        int
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

func NewSkill(name string, source string, url string, size string) *SkillIcon {
	skill := SkillIcon{
		Name:   name,
		ImgURL: fmt.Sprintf("fa-%s fa-%s fa-%s", source, url, size),
	}
	return &skill
}

func NewProject(title string, description string, imgURL string, skills [][]string, link string, typeProject int) *ProjectHero {
	project := ProjectHero{
		Title:       title,
		Description: description,
		ImgURL:      imgURL,
		Skills:      skills,
		Link:        link,
		Type:        typeProject,
	}
	return &project
}
