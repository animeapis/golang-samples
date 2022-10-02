package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/type/localized_text"

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

	request := &knowledge.ReviewContributionRequest{
		Name:    ContributionParent,
		Comment: "missing esclamation marks",
		Changes: &knowledge.ContributionChanges{
			Additions: []*knowledge.EntryEntity{
				{
					Entity: &knowledge.EntryEntity_Anime{
						Anime: &knowledge.Anime{
							Name: "animes/4389926960196515999",
							Aliases: []*localized_text.LocalizedText{
								{
									LanguageCode: "ita",
									Text:         "All'Arrembaggio!!!",
								},
							},
						},
					},
				},
			},
		},
		Discards: &knowledge.ContributionChanges{
			Additions: []*knowledge.EntryEntity{
				{
					Entity: &knowledge.EntryEntity_Anime{
						Anime: &knowledge.Anime{
							Name: "animes/4389926960196515999",
							Aliases: []*localized_text.LocalizedText{
								{
									LanguageCode: "ita",
									Text:         "All'Arrembaggio",
								},
							},
						},
					},
				},
			},
		},
	}

	contribution, err := client.ReviewContribution(ctx, request)
	if nil != err {
		log.Fatalf("ReviewContribution: %s", err)
	}

	log.Println("---------------------------------------------------------")
	log.Printf("name        : %s", contribution.GetName())
	log.Printf("displayName : %s", contribution.GetDisplayName())
	log.Printf("state       : %s", contribution.GetState())
	log.Printf("reviewer 	: %s", contribution.GetReviewer())
	log.Printf("generation 	: %d", contribution.GetGeneration())
}
