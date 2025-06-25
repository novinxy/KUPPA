package main

import (
	"fmt"
	"html/template"
	"os"
	"os/exec"
	"strings"
)

type commit struct {
	CommitId      string
	Author        string
	Date          string
	Title         string
	PullRequestId string
	Message       string
}

func getGitCommits(repo string) []commit {
	format := `[[start]]%n%H%n%al%n%aI%n%s%n%b`
	cmd := fmt.Sprintf("cd %s && git log --branches=*master --after=2025-04-01 --author=grno --format=format:%s", repo, format)

	gitLogCmd := exec.Command("cmd", "/c", cmd)
	out, err := gitLogCmd.Output()
	if err != nil {
		panic(fmt.Sprintf("Failed on running %s", repo))
	}

	data := string(out)
	commits := strings.Split(data, "[[start]]")
	cleanCommits := []commit{}

	for _, c := range commits[1:] {
		commitLines := strings.Split(strings.TrimSpace(c), "\n")

		subject := commitLines[3]
		subjectParts := strings.Split(subject, ":")

		commitData := commit{
			CommitId:      commitLines[0],
			Author:        commitLines[1],
			Date:          commitLines[2],
			PullRequestId: strings.TrimSpace(strings.ReplaceAll(subjectParts[0], "Merged PR ", "")),
			Title:         strings.TrimSpace(subjectParts[1]),
			Message:       strings.Join(commitLines[4:], "\n"),
		}

		cleanCommits = append(cleanCommits, commitData)
	}
	return cleanCommits
}

var repositories = []string{
	"C:\\Git\\JenkinsJobs",
	"C:\\Git\\TestComplete",
}

func main() {

	allCommits := []commit{}

	for _, repo := range repositories {
		commits := getGitCommits(repo)
		allCommits = append(allCommits, commits...)
	}

	// funcs := template.FuncMap{
	// 	"nl2br": func(text string) template.HTML {
	// 		escaped := template.HTMLEscapeString(text)
	// 		return template.HTML(strings.ReplaceAll(escaped, "\n", "<br>"))
	// 	},
	// }
	t := template.New("report.tpl")
	// t.Funcs(funcs)
	t, err := t.ParseFiles("report.tpl")

	if err != nil {
		panic(fmt.Sprintf("Failed on template creation with %v", err))
	}

	e := t.Execute(os.Stdout, allCommits)
	if e != nil {
		panic(e)
	}

}
