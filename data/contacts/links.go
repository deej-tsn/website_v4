package contacts

import "github.com/deej-tsn/website_v4/internal/models"

var Links []models.ContactIcon
var pathToIcons string = "static/assets/contactIcons"

var github *models.ContactIcon = models.NewContact("Github", pathToIcons, "https://github.com/deej-tsn")
var linkedIn *models.ContactIcon = models.NewContact("LinkedIn", pathToIcons, "https://www.linkedin.com/in/dempsey-palacio-tascon/")
var email *models.ContactIcon = models.NewContact("Email", pathToIcons, "mailto: dempsey.2001@hotmail.co.uk")
var cv *models.ContactIcon = models.NewContact("CV", pathToIcons, "static/assets/documents/Dempsey's CV - Fullstack.pdf")

func GetLinks() []models.ContactIcon {
	Links = append(Links, *github, *linkedIn, *email, *cv)
	return Links
}
