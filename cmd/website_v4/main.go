package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	"github.com/deej-tsn/website_v4/internal/helper"
	"github.com/deej-tsn/website_v4/internal/routes"

	layoutComponents "github.com/deej-tsn/website_v4/web/components/layout"
	blogComponents "github.com/deej-tsn/website_v4/web/components/pages/blog"
	homeComponents "github.com/deej-tsn/website_v4/web/components/pages/home"
	projectComponents "github.com/deej-tsn/website_v4/web/components/pages/project"
)

func main() {

	//db := sql.ConnectToDatabase()
	e := echo.New()

	pathToPosts := "./data/blog/posts/"
	pathToWeb := "./web/public"

	//CONTROLLERS

	blogController := routes.NewSlugReader(pathToPosts)

	// MIDDLEWARE

	// logs all http requests
	e.Use(middleware.Logger())

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":80"
	}
	e.Server.Addr = PORT

	// static files in public folder
	e.Static("static", pathToWeb)

	// resets server if error
	e.Use(middleware.Recover())

	// CORS - cross origin resource sharing
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// allow any origin - DEVELOPMENT ONLY
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	home := layoutComponents.Index("Home", 0, homeComponents.HomeHub())
	projectHome := layoutComponents.Index("Project", 1, projectComponents.Content())
	blogHome := layoutComponents.Index("Blog", 2, blogComponents.Content(blogComponents.LoadPostList()))

	// ROUTES

	///HOME
	e.GET("/", func(c echo.Context) error {
		return helper.Render(c, http.StatusOK, home)
	})

	//PROJECTS
	e.GET("/projects", func(c echo.Context) error {
		return helper.Render(c, http.StatusOK, projectHome)
	})
	e.GET("/projects/posts", routes.GetProjects)

	/// BLOG
	e.GET("/blog", func(c echo.Context) error {
		return helper.Render(c, http.StatusOK, blogHome)
	})
	e.GET("/blog/posts/:slug", blogController.GetPost)
	e.GET("/blog/posts", blogController.GetAllPosts)

	e.Logger.Fatal(e.Start(":8080"))
}
