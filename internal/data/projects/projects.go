package projects

import (
	"github.com/deej-tsn/website_v4/internal/helper"
	"github.com/deej-tsn/website_v4/internal/models"
)

var Projects []models.ProjectHero = []models.ProjectHero{

	*models.NewProject(
		"Split bills with a snap",
		"",
		"static/assets/projects/billsplitter.webp",
		[][]string{{""}},
		"https://billsplitter.dempseypalaciotascon.com",
		helper.PROJECTS_FRONTEND,
	),

	*models.NewProject(
		"collabarative Sudoku (with Chat)",
		"",
		"static/assets/projects/compSudoku.webp",
		[][]string{{""}},
		"https://sudoku.dempseypalaciotascon.com",
		helper.PROJECTS_BACKEND,
	),

	*models.NewProject(
		"freeCodeCamp London Website",
		"Crafted a clean, modern UI with a focus on simplicity and intuitive navigation, ensuring seamless integration of community links.",
		"static/assets/projects/freeCodeCampLdn.webp",
		[][]string{{"brands", "javascript", "js"}},
		"https://freecodecamp.london/",
		helper.PROJECTS_FRONTEND,
	),

	*models.NewProject(
		"Wordle",
		"Building a replica of wordle was fun and enhanced my skills. It provides hands-on experience in React Components design.",
		"static/assets/projects/Wordle.webp",
		[][]string{{"brands", "react", "react"}, {"brands", "javascript", "js"}},
		"https://wordleclonedempsey.netlify.app",
		helper.PROJECTS_FRONTEND,
	),
	*models.NewProject(
		"Color Theory",
		"See how good your ability to match colours is!",
		"static/assets/projects/ColourGame.webp",
		[][]string{{"brands", "react", "react"}, {"brands", "javascript", "js"}},
		"https://colourgamedempsey.netlify.app/",
		helper.PROJECTS_FRONTEND,
	),

	*models.NewProject(
		"User Authentication",
		"Fullstack app focused on robust user authentication, exploring modern techniques like JWTs and security.",
		"static/assets/projects/UserAuth.webp",
		[][]string{
			{"brands", "react", "react"},
			{"brands", "javaScript", "js"},
			{"brands", "Golang", "golang"},
			{"solid", "MySQL", "database"},
		},
		"https://github.com/deej-tsn/user-auth",
		helper.PROJECTS_BACKEND,
	),
}
