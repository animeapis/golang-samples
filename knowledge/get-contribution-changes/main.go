package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/knowledge/v1alpha1"
	knowledge "github.com/animeapis/go-genproto/knowledge/v1alpha1"
)

var (
	ContributionName = "[CONTRIBUTION-NAME]"

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

	request := &knowledge.GetContributionChangesRequest{
		Name: ContributionName,
	}

	changes, err := client.GetContributionChanges(ctx, request)
	if nil != err {
		log.Fatalf("GetContributionChanges: %s", err)
	}

	log.Println("---------------------------------------------------------")
	for i, addition := range changes.Additions {
		anime := addition.Entity.(*knowledge.EntryEntity_Anime).Anime
		log.Printf("add[%d] name        : %s", i, anime.GetName())
	}
	for i, deletion := range changes.Deletions {
		anime := deletion.Entity.(*knowledge.EntryEntity_Anime).Anime
		log.Printf("del[%d] name        : %s", i, anime.GetName())
	}
}
