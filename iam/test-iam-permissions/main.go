package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/iam/v1"

	gapic "github.com/animeapis/api-go-client/image/v1alpha1"
)

var (
	ServiceAccount = "[SERVICE-ACCOUNT]"

	TestPermissions = []string{
		"iam.serviceAccounts.get",
		"iam.serviceAccounts.create",
		"iam.serviceAccounts.setIamPolicy",
	}

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "image.animeapis.com:443"
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
		Resource:    ServiceAccount,
		Permissions: TestPermissions,
	}

	response, err := client.TestIamPermissions(ctx, request)
	if err != nil {
		log.Fatalf("TestIamPermissions: %s", err)
	}

	log.Printf("resource   				 : %s", ServiceAccount)
	log.Printf("allowed permissions: %v", response.GetPermissions())
}
