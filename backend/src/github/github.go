package github

import (
	"context"
	"github.com/TheoEb/scalingo_test/backend/src/models"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"sync"
)

type Client struct {
	Client *github.Client
}

func NewClient(apiKey string) *Client {
	ctx := context.Background()
	tokenSource := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: apiKey,
		},
	)
	oauthClient := oauth2.NewClient(ctx, tokenSource)
	client := github.NewClient(oauthClient)
	return &Client{
		Client: client,
	}
}

func (c *Client) ListRepositories() ([]*github.Repository, error) {
	ctx := context.Background()
	repos, _, err := c.Client.Repositories.ListAll(ctx, nil)
	if err != nil {
		return nil, err
	}
	return repos, nil
}

func (c *Client) GetLanguageAndLines(repos []*github.Repository) ([]*models.Data, error) {
	ctx := context.Background()
	wg := sync.WaitGroup{}
	mx := sync.Mutex{}

	var data []*models.Data
	for _, repo := range repos {
		wg.Add(1)
		go func(repo *github.Repository) {
			defer wg.Done()
			repoData := &models.Data{
				Name: repo.GetFullName(),
				URL:  repo.GetHTMLURL(),
			}
			lgs, _, err := c.Client.Repositories.ListLanguages(ctx, repo.GetOwner().GetLogin(), repo.GetName())
			if err != nil {
				return
			}
			for lg, lines := range lgs {
				repoData.Language = append(repoData.Language, lg)
				repoData.Lines = append(repoData.Lines, lines)
			}
			mx.Lock()
			data = append(data, repoData)
			mx.Unlock()
		}(repo)
	}
	wg.Wait()
	return data, nil
}
