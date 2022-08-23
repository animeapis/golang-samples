package main

import (
	"context"
	"log"
	"strings"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/genproto/googleapis/iam/v1"

	gapic "github.com/animeapis/api-go-client/identity/v1alpha1"
)

var (
	User   = "[USER]"
	Member = "[MEMBER]"

	Role = "roles/owner"

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

	request := &iam.SetIamPolicyRequest{
		Resource: User,
		Policy: &iam.Policy{
			Version: 1,
			Bindings: []*iam.Binding{
				{
					Role: Role,
					Members: []string{
						"user:" + strings.TrimPrefix(User, "users/"),
						Member,
					},
				},
			},
		},
	}

	policy, err := client.SetIamPolicy(ctx, request)
	if err != nil {
		log.Fatalf("SetIamPolicy: %s", err)
	}

	log.Printf("resource: %s", User)
	log.Printf("bindings: %v", policy.Bindings)
}
