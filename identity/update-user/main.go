package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	gapic "github.com/animeapis/api-go-client/identity/v1alpha1"
	identity "github.com/animeapis/go-genproto/identity/v1alpha1"
)

var (
	User     = "[USER]"
	Username = "[USERNAME]"

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

	request := &identity.UpdateUserRequest{
		User: &identity.User{
			Name:     User,
			Username: Username,
		},
		UpdateMask: &fieldmaskpb.FieldMask{
			Paths: []string{
				"username", // Partial update: modify only the username.
			},
		},
	}

	response, err := client.UpdateUser(ctx, request)
	if err != nil {
		log.Fatalf("UpdateUser: %s", err)
	}

	log.Printf("user         : %s", User)
	log.Printf("username     : %s", response.Username)
	log.Printf("discriminator: %s", response.Discriminator)
}
