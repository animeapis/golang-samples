package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/image/v1alpha1"
	image "github.com/animeapis/go-genproto/image/v1alpha1"
)

var (
	Image = "[IMAGE]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "image.animeapis.com:443"
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

	client, err := gapic.NewImageRouterClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewClient: %s", err)
	}

	request := &image.GetImageRouteRequest{
		Name: Image,
	}

	response, err := client.GetImageRoute(ctx, request)
	if err != nil {
		log.Fatalf("GetImageRoute: %s", err)
	}

	log.Printf("[cdn] public url: %s", response.GetUrl())
}
