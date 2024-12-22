package routes

import (
	"log"
	"net/http"

	"github.com/deej-tsn/website_v4/internal/data/projects"
	"github.com/deej-tsn/website_v4/internal/helper"
	"github.com/deej-tsn/website_v4/internal/models"
	projectComponents "github.com/deej-tsn/website_v4/web/components/pages/project"
	"github.com/labstack/echo/v4"
)

func GetProjects(c echo.Context) error {
	projects := projects.Projects
	filter := filterToInt(c.QueryParam("filter"))
	projects = filterProjects(projects, filter)
	return helper.Render(c, http.StatusAccepted, projectComponents.ProjectList(projects))
}

func filterToInt(filterString string) int {

	var filter int
	switch filterString {
	case "all":
		filter = helper.PROJECTS_ALL
	case "frontend":
		filter = helper.PROJECTS_FRONTEND
	case "backend":
		filter = helper.PROJECTS_BACKEND
	default:
		log.Println("illegal filter")
		filter = helper.PROJECTS_ALL
	}
	return filter
}

func filterProjects(projects []models.ProjectHero, typeProjects int) []models.ProjectHero {
	if typeProjects == helper.PROJECTS_ALL {
		return projects
	}
	if typeProjects == helper.PROJECTS_FRONTEND {
		return helper.Filter(projects, isFrontend)
	}
	return helper.Filter(projects, isBackend)
}

func isValidType(project models.ProjectHero, typeProjects int) bool {
	return (project.Type == typeProjects)
}

func isFrontend(project models.ProjectHero) bool {
	return isValidType(project, helper.PROJECTS_FRONTEND)
}

func isBackend(projects models.ProjectHero) bool {
	return isValidType(projects, helper.PROJECTS_BACKEND)
}
