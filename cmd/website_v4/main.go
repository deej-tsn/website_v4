package main

import (
	"github.com/labstack/echo/v4"

	"github.com/labstack/echo/v4/middleware"

	"github.com/deej-tsn/website_v4/internal/helper"
	"github.com/deej-tsn/website_v4/internal/routes"
	layoutComponents "github.com/deej-tsn/website_v4/web/components/layout"
	blogComponents "github.com/deej-tsn/website_v4/web/components/pages/blog"
	homeComponents "github.com/deej-tsn/website_v4/web/components/pages/home"
)

func main() {

	//db := sql.ConnectToDatabase()
	e := echo.New()

	pathToData := "./data/posts/"
	pathToWeb := "./web/public"

	//CONTROLLERS

	blogController := routes.NewSlugReader(pathToData)

	// MIDDLEWARE

	// logs all http requests
	e.Use(middleware.Logger())

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

	index := layoutComponents.Index("Home", 0, homeComponents.HomeHub())
	blogHome := layoutComponents.Index("Blog", 2, blogComponents.Content(blogComponents.LoadPostList()))

	// ROUTES
	e.GET("/", func(c echo.Context) error {
		return helper.Render(c, 100, index)
	})

	e.GET("/blog", func(c echo.Context) error {
		return helper.Render(c, 100, blogHome)
	})
	e.GET("/blog/posts/:slug", blogController.GetPost)
	e.GET("/blog/posts", blogController.GetAllPosts)

	e.Logger.Fatal(e.Start(":8080"))
}
