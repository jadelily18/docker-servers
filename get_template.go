package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/google/go-github/v47/github"
)

func get_template_names(client *github.Client) []string {
	_, dockerComposeDirContent, _, err := client.Repositories.GetContents(context.Background(), "jadelily18", "docker-servers", "docker-compose", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var template_names []string
	for i := 0; i < len(dockerComposeDirContent); i++ {
		template_names = append(template_names, dockerComposeDirContent[i].GetName())
	}

	return template_names
}

func get_download_url(client *github.Client, template_name string) string {
	fileContent, _, _, err := client.Repositories.GetContents(context.Background(), "jadelily18", "docker-servers", "docker-compose/"+template_name+"/docker-compose.yml", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fileContent.GetDownloadURL()
}

func download_file(url string, new_dir_name string) {

	os.Mkdir(new_dir_name, os.ModeDir)
	out, err := os.Create(filepath.Join(new_dir_name, "docker-compose.yml"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	client := github.NewClient(nil)

	var template_names = get_template_names(client)

	// the answers will be written to this struct
	answers := struct {
		ServerDir    string `survey:"server_directory"`
		TemplateName string `survey:"template_name"`
	}{}

	var questions = []*survey.Question{
		{
			Name:     "server_directory",
			Prompt:   &survey.Input{Message: "What do you want to name your server directory?", Default: "mc-server"},
			Validate: survey.Required,
		},
		{
			Name:     "template_name",
			Prompt:   &survey.Select{Message: "Which template do you want to use?", Options: template_names},
			Validate: survey.Required,
		},
	}

	// ask the questions
	survey_ask_err := survey.Ask(questions, &answers)
	if survey_ask_err != nil {
		fmt.Println(survey_ask_err.Error())
		os.Exit(1)
	}

	if _, err := os.Stat(answers.ServerDir); !os.IsNotExist(err) {
		fmt.Printf("\nDirectory '%s' already exists! Please retry with a different server directory name.\n", answers.ServerDir)
	} else {
		var downloadUrl = get_download_url(client, answers.TemplateName)
		download_file(downloadUrl, answers.ServerDir)
		fmt.Println("\nYou can now run your server using `docker compose up` in the newly-created directory. If you don't already have docker installed, you can use its their convenience script at https://github.com/docker/docker-install.\n")
	}

}
