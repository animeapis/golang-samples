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

	Resource1 = "[RESOURCE-1]"
	Resource2 = "[RESOURCE-2]"
	Resource3 = "[RESOURCE-3]"

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

	request := &library.BatchCreatePlaylistItemsRequest{
		Parent: Playlist,
		Items: []*library.PlaylistItem{
			{Resource: Resource1},
			{Resource: Resource2},
			{Resource: Resource3},
		},
	}

	response, err := client.BatchCreatePlaylistItems(ctx, request)
	if err != nil {
		log.Fatalf("BatchCreatePlaylistItems: %s", err)
	}

	for _, item := range response.Items {
		log.Println("---------------------------------------------------------")
		log.Printf("name       : %s", item.GetName())
		log.Printf("resource   : %s", item.GetResource())
		log.Printf("create time: %s", item.GetCreateTime().AsTime().String())
	}
}
