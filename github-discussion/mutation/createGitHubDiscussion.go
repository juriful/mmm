package mutation

import (
	"context"
	"fmt"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
	"os"
)

var token = os.Getenv("GITHUB_TOKEN")
var endpoint = "https://api.github.com/graphql"

func CreateGitHubDiscussion() error {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client := githubv4.NewEnterpriseClient(endpoint, httpClient)


	var m struct {
		CreateDiscussion struct {
			Discussion struct{
				Number githubv4.Int
			}
		} `graphql:"createDiscussion(input: $input)"`
	}
	input := githubv4.CreateDiscussionInput{
		RepositoryID:     "R_kgDOGS_ajg",
		Title:            "Titleテスト from Golang!",
		Body:             "Bodyテスト",
		CategoryID:       "DIC_kwDOGS_ajs4B_qUe",
	}

	err := client.Mutate(context.Background(), &m, input, nil)
	if err != nil {
		return err
	}

	fmt.Println("Added discussion: ", m.CreateDiscussion.Discussion.Number)
	return nil
}
