package main

import (
	"fmt"
	"pull-request-formatter/pkg/config"
	"pull-request-formatter/pkg/git"
	"pull-request-formatter/pkg/log"
	"regexp"
	"strings"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Error(err)
		return
	}

	//prompt, err := getPrompt()
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//changelog, err := openai.Send(prompt)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	//
	//log.SaveToFile(changelog, "changelog")
	//
	//err = git.SetPRDescription(changelog)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}

	err = createPrComment()
	if err != nil {
		log.Error(err)
		return
	}

	log.Success()
}

func getPrompt() (prompt string, err error) {
	commits, err := git.GetCommits()
	if err != nil {
		return
	}

	prompt = config.PromptPreText

	for _, c := range commits {
		message := c.Commit.Message

		//check if the commit contains a changelog message
		if !strings.HasPrefix(message, "#changelog") {
			continue
		}

		message = strings.Replace(message, "#changelog ", "", 1)

		// remove links from the message
		re := regexp.MustCompile(`\bhttps?://\S+`)
		message = re.ReplaceAllString(message, "")

		// remove all newlines from the message
		message = strings.ReplaceAll(message, "\n", "")

		prompt += "\n - " + message
	}

	prompt += "\n" + config.PromptAfterText

	log.SaveToFile(prompt, "prompt")

	return
}

func createPrComment() (err error) {
	versions, err := git.GetVersions()
	if err != nil {
		return
	}

	//if there are no changes present, the versions array will be empty, therefore, no need to add a pr comment
	if len(versions) == 0 {
		return
	}

	versionsLogText := config.VersionsLogPreText

	for _, versionObj := range versions {
		versionsLogText += fmt.Sprintf("| %s | %s | %d |\n", versionObj.Name, versionObj.OldVersion, versionObj.OldIntVersion)
	}

	versionsLogText += config.VersionsLogPostText

	for _, versionObj := range versions {
		versionsLogText += fmt.Sprintf("| %s | %s | %d |\n", versionObj.Name, versionObj.NewVersion, versionObj.NewIntVersion)
	}

	return git.CreatePRComment(versionsLogText)
}
