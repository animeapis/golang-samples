package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/library/v1alpha1"
	library "github.com/animeapis/go-genproto/library/v1alpha1"
)

var (
	Playlist = "[PLAYLIST]"
	Resource = "[RESOURCE]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "library.animeapis.com:443"
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

	request := &library.CreatePlaylistItemRequest{
		Parent: Playlist,
		Item: &library.PlaylistItem{
			Resource: Resource,
		},
	}

	item, err := client.CreatePlaylistItem(ctx, request)
	if err != nil {
		log.Fatalf("CreatePlaylistItem: %s", err)
	}

	log.Printf("name       : %s", item.Name)
	log.Printf("resource   : %s", item.Resource)
	log.Printf("create time: %s", item.CreateTime.AsTime().String())
}
