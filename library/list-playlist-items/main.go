package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/library/v1alpha1"
	library "github.com/animeapis/go-genproto/library/v1alpha1"
)

var (
	Playlist = "[PLAYLIST]"

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

	request := &library.ListPlaylistItemsRequest{
		Parent: Playlist,
	}

	it := client.ListPlaylistItems(ctx, request)
	for {
		item, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				return
			}

			log.Fatalf("ListPlaylistItems: %s", err)
		}

		log.Println("---------------------------------------------------------")
		log.Printf("name       : %s", item.GetName())
		log.Printf("resource   : %s", item.GetResource())
		log.Printf("create time: %s", item.GetCreateTime().AsTime().String())
	}
}
