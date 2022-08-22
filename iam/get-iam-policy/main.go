package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/iam/v1"

	gapic "github.com/animeapis/api-go-client/iam/admin/v1alpha1"
)

var (
	ServiceAccount = "[SERVICE-ACCOUNT]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "iam.animeapis.com:443"
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

	client, err := gapic.NewIamClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewIamClient: %s", err)
	}

	request := &iam.GetIamPolicyRequest{
		Resource: ServiceAccount,
	}

	policy, err := client.GetIamPolicy(ctx, request)
	if err != nil {
		log.Fatalf("GetIamPolicy: %s", err)
	}

	log.Printf("resource: %s", ServiceAccount)
	log.Printf("bindings: %v", policy.Bindings)
}
