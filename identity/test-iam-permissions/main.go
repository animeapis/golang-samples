package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/iam/v1"

	gapic "github.com/animeapis/api-go-client/identity/v1alpha1"
)

var (
	User = "[USER]"

	TestPermissions = []string{
		"identity.users.get",
		"identity.users.getSettings",
		"identity.users.setIamPolicy",
	}

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

	request := &iam.TestIamPermissionsRequest{
		Resource:    User,
		Permissions: TestPermissions,
	}

	response, err := client.TestIamPermissions(ctx, request)
	if err != nil {
		log.Fatalf("TestIamPermissions: %s", err)
	}

	log.Printf("resource           : %s", User)
	log.Printf("allowed permissions: %v", response.GetPermissions())
}
