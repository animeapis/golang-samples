package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/iam/v1"

	gapic "github.com/animeapis/api-go-client/library/v1alpha1"
)

var (
	Playlist = "[PLAYLIST]"

	Member = "allUsers"
	Role   = "roles/library.viewer"

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

	request := &iam.SetIamPolicyRequest{
		Resource: Playlist,
		Policy: &iam.Policy{
			Version: 1,
			Bindings: []*iam.Binding{
				{
					Role:    Role,
					Members: []string{Member},
				},
			},
		},
	}

	policy, err := client.SetIamPolicy(ctx, request)
	if err != nil {
		log.Fatalf("SetIamPolicy: %s", err)
	}

	log.Printf("resource: %s", Playlist)
	log.Printf("bindings: %v", policy.Bindings)
}
