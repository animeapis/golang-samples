package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/knowledge/v1alpha1"
	knowledge "github.com/animeapis/go-genproto/knowledge/v1alpha1"
)

var (
	ContributionParent = "[USER-OR-ORGANIZATION]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "knowledge.animeapis.com:443"
)

func main() {
	ctx := context.Background()

	config := &clientcredentials.Config{
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		TokenURL:     TokenURL,
	}

	options := []option.ClientOption{
		option.WithEndpoint(Endpoint),
		option.WithTokenSource(config.TokenSource(ctx)),
	}

	client, err := gapic.NewClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewClient: %s", err)
	}

	request := &knowledge.ListContributionsRequest{
		Parent: ContributionParent,
	}

	it := client.ListContributions(ctx, request)
	for {
		contribution, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				log.Println("---------------------------------------------------------")
				return
			}

			log.Fatalf("ListContributions: %s", err)
		}

		log.Println("---------------------------------------------------------")
		log.Printf("name        : %s", contribution.GetName())
	}
}
