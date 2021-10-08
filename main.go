package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/urfave/cli/v2"
)

func main() {
	app := configCLI(run)
	err := app.Run([]string{""})
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) error {
	terms := readSearchTerms()

	url := "https://github.com/search?o=desc&q=" + terms + "&s=stars&type=Repositories"
	fmt.Println("Search:")
	fmt.Println("\t", url)

	resp, err := soup.Get(url)
	if err != nil {
		return errors.New("There was a problem during the search.")
	}

	doc := soup.HTMLParse(resp)
	repos := doc.FindAll("li", "class", "repo-list-item")

	fmt.Println("Result:")
	for i, repo := range repos {
		raw_rank := ("   " + strconv.Itoa(i+1))
		rank := raw_rank[len(raw_rank)-2:]
		name := "https://github.com/" + repo.Find("a", "class", "v-align-middle").FullText()
		star := (strings.TrimSpace(repo.Find("a", "class", "Link--muted").Text()) + "   ")[:5]
		fmt.Println("\t", rank, ":", star, name)
	}
	return nil
}

func readSearchTerms() string {
	fmt.Print("🔍 > ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	target := strings.Replace(line, " ", "+", -1)
	return target
}

func configCLI(f func(c *cli.Context) error) *cli.App {
	cli.AppHelpTemplate = `Name:
	gh-topstar - interactive(?) search Top 10 GitHub Repogitories ordered by star

Usage:
	$ gh-topstar

Example:
`

	app := &cli.App{
		Name:  "gh-topstar",
		Usage: "interactive(?) search Top 10 GitHub Repogitories ordered by star",
		Action: func(c *cli.Context) error {
			return f(c)
		},
	}
	return app
}
