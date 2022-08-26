package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/credentials/v1alpha1"
	credentials "github.com/animeapis/go-genproto/credentials/v1alpha1"
)

var (
	Credentials = "[CREDENTIALS]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "credentials.animeapis.com:443"
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

	client, err := gapic.NewKeeperClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewKeeperClient: %s", err)
	}

	request := &credentials.ActAsCredentialsRequest{
		Name: Credentials,
	}

	response, err := client.ActAsCredentials(ctx, request)
	if err != nil {
		log.Fatalf("ActAsCredentials: %s", err)
	}

	log.Printf("[oauth2] access token: %s", response.AccessToken)

	log.Printf("[basic] username     : %s", response.Basic.GetUsername())
	log.Printf("[basic] password     : %s", response.Basic.GetPassword())
}
