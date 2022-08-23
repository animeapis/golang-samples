package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/identity/v1alpha1"
	identity "github.com/animeapis/go-genproto/identity/v1alpha1"
)

var (
	User = "[USER]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "identity.animeapis.com:443"
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

	request := &identity.GetUserDefaultsRequest{
		Name: User,
	}

	response, err := client.GetUserDefaults(ctx, request)
	if err != nil {
		log.Fatalf("GetUserDefaults: %s", err)
	}

	log.Printf("user          : %s", User)
	log.Printf("album banner  : %s", response.AlbumBanner)
	log.Printf("album profile : %s", response.AlbumProfile)
	log.Printf("playlist later: %s", response.PlaylistLater)
	log.Printf("playlist liked: %s", response.PlaylistLiked)
}
