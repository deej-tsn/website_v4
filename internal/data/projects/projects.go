package projects

import "github.com/deej-tsn/website_v4/internal/models"

var Projects []models.ProjectHero = []models.ProjectHero{

	*models.NewProject(
		"Wordle",
		"A clone of Wordle",
		"static/assets/projects/WordleClone.png",
		[][]string{{"react", "react"}, {"javascript", "js"}},
		"Frontend",
		"https://wordleclonedempsey.netlify.app",
	),
	*models.NewProject(
		"Color Theory",
		"See how good your ability to match colours is!",
		"static/assets/projects/ColourGame.png",
		[][]string{{"react", "react"}, {"javascript", "js"}},
		"Frontend",
		"https://wordleclonedempsey.netlify.app",
	),
}
