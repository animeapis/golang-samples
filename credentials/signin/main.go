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
	Flow = "[FLOW]"

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

	client, err := gapic.NewOAuth2Client(ctx, options...)
	if err != nil {
		log.Fatalf("NewOAuth2Client: %s", err)
	}

	request := &credentials.SignInRequest{
		Name: Flow,
	}

	response, err := client.SignIn(ctx, request)
	if err != nil {
		log.Fatalf("SignIn: %s", err)
	}

	log.Printf("authorization url: %s", response.AuthorizationUrl)
}
