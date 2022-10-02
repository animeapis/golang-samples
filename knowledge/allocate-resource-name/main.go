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

	request := &knowledge.AllocateResourceNameRequest{
		Kind: "Anime",
	}

	name, err := client.AllocateResourceName(ctx, request)
	if err != nil {
		log.Fatalf("AllocateResourceName: %s", err)
	}

	log.Println("---------------------------------------------------------")
	log.Printf("name        : %s", name)
}
