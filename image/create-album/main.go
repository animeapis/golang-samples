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
	AlbumParent      = "[USER-OR-ORGANIZATION]"
	AlbumDisplayName = "[DISPLAY-NAME]"

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

	client, err := gapic.NewClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewClient: %s", err)
	}

	request := &image.CreateAlbumRequest{
		Parent: AlbumParent,
		Album: &image.Album{
			DisplayName: AlbumDisplayName,
		},
	}

	album, err := client.CreateAlbum(ctx, request)
	if err != nil {
		log.Fatalf("CreateAlbum: %s", err)
	}

	log.Printf("name        : %s", album.GetName())
	log.Printf("display name: %s", album.GetDisplayName())
}
