package projects

import (
	"github.com/deej-tsn/website_v4/internal/helper"
	"github.com/deej-tsn/website_v4/internal/models"
)

var Projects []models.ProjectHero = []models.ProjectHero{

	*models.NewProject(
		"freeCodeCamp London Website",
		"Crafted a clean, modern UI with a focus on simplicity and intuitive navigation, ensuring seamless integration of community links.",
		"static/assets/projects/freeCodeCamp.png",
		"cover",
		[][]string{{"brands", "javascript", "js"}},
		"https://freecodecamp.london/",
		helper.PROJECT_FRONTEND,
	),

	*models.NewProject(
		"Wordle",
		"Building a replica of wordle was fun and enhanced my skills. It provides hands-on experience in React Components design.",
		"static/assets/projects/WordleClone.png",
		"cover",
		[][]string{{"brands", "react", "react"}, {"brands", "javascript", "js"}},
		"https://wordleclonedempsey.netlify.app",
		helper.PROJECT_FRONTEND,
	),
	*models.NewProject(
		"Color Theory",
		"See how good your ability to match colours is!",
		"static/assets/projects/ColourGame.png",
		"contain",
		[][]string{{"brands", "react", "react"}, {"brands", "javascript", "js"}},
		"https://colourgamedempsey.netlify.app/",
		helper.PROJECT_FRONTEND,
	),

	*models.NewProject(
		"User Authentication",
		"Fullstack app focused on robust user authentication, exploring modern techniques like JWTs and security.",
		"static/assets/projects/UserAuth.svg",
		"contain",
		[][]string{
			{"brands", "react", "react"},
			{"brands", "javaScript", "js"},
			{"brands", "Golang", "golang"},
			{"solid", "MySQL", "database"},
		},
		"https://github.com/deej-tsn/user-auth",
		helper.PROJECT_BACKEND,
	),
}
