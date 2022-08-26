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
	PlaylistParent      = "[USER-OR-ORGANIZATION]"
	PlaylistDisplayName = "[DISPLAY-NAME]"

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

	request := &library.CreatePlaylistRequest{
		Parent: PlaylistParent,
		Playlist: &library.Playlist{
			DisplayName: PlaylistDisplayName,
			Type:        library.Type_CUSTOM,
		},
	}

	playlist, err := client.CreatePlaylist(ctx, request)
	if err != nil {
		log.Fatalf("CreatePlaylist: %s", err)
	}

	log.Printf("name        : %s", playlist.Name)
	log.Printf("display name: %s", playlist.DisplayName)
	log.Printf("type        : %s", playlist.Type.String())
}
