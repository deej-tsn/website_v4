package skills

import "github.com/deej-tsn/website_v4/internal/models"

var skills [][]string = [][]string{{"javascript", "js"}, {"react", "react"}, {"java", "java"}, {"docker", "docker"}, {"golang", "golang"}, {"CSS", "css"}, {"html", "html5"}, {"Python", "python"}}

func GenerateSkillsArray(skills [][]string) []models.SkillIcon {
	skillIcons := []models.SkillIcon{}
	for _, skill := range skills {
		skillIcons = append(skillIcons, *models.NewSkill(skill[0], skill[1]))
	}
	return skillIcons
}

var SkillIcons []models.SkillIcon = GenerateSkillsArray(skills)
