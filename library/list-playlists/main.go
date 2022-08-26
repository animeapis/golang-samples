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
	PlaylistParent = "[USER-OR-ORGANIZATION]"

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

	request := &library.ListPlaylistsRequest{
		Parent: PlaylistParent,
	}

	it := client.ListPlaylists(ctx, request)
	for {
		playlist, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				return
			}

			log.Fatalf("ListPlaylists: %s", err)
		}

		log.Println("---------------------------------------------------------")
		log.Printf("name        : %s", playlist.GetName())
		log.Printf("display name: %s", playlist.GetDisplayName())
		log.Printf("type        : %s", playlist.GetType())
	}
}
