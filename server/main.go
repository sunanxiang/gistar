package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	personalAccessToken = "093eea479dc6207c52c41b7b8b82aa1dd0b93335"
)

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func main() {
	tokenSource := &TokenSource{
		AccessToken: personalAccessToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := github.NewClient(oauthClient)

	user, _, err := client.Users.Get(context.Background(), "")
	if err != nil {
		fmt.Printf("client.Users.Get() faled with '%s'\n", err)
		return
	}

	d, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		fmt.Printf("json.MarshlIndent() failed with %s\n", err)
		return
	}

	fmt.Printf("User:\n%s\n", string(d))
}