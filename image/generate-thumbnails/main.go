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

	log.Printf("[image #1]")
	log.Printf("  - config: original")
	log.Printf("  - url   : %s", response.GetUrl())

	log.Printf("[image #2]")
	log.Printf("  - config: max width equal to 256px")
	log.Printf("  - url   : %s=w256", response.GetUrl())

	log.Printf("[image #3]")
	log.Printf("  - config: max height equal to 144px")
	log.Printf("  - url   : %s=h144", response.GetUrl())

	log.Printf("[image #4]")
	log.Printf("  - config: max width equal to 144px and max width equal to 144px")
	log.Printf("  - url   : %s=s144", response.GetUrl())
	log.Printf("  - alias : %s=w144-h144", response.GetUrl())

	log.Printf("[image #6]")
	log.Printf("  - config: convert to WebP image format")
	log.Printf("  - url   : %s=webp", response.GetUrl())

	log.Printf("[image #7]")
	log.Printf("  - config: convert to PNG image format")
	log.Printf("  - url   : %s=png", response.GetUrl())

	log.Printf("[image #8]")
	log.Printf("  - config: convert to JPEG image format")
	log.Printf("  - url   : %s=jpeg", response.GetUrl())

	log.Printf("[image #9]")
	log.Printf("  - config: convert to WebP image format and max width 120px")
	log.Printf("  - url   : %s=w120-webp", response.GetUrl())
}
