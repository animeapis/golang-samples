package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/iam/admin/v1alpha1"
	iam "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
)

var (
	ServiceAccountName        = "[NAME]"
	ServiceAccountDisplayName = "[DISPLAY-NAME]"
	ServiceAccountDescription = "[DESCRIPTION]"

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

	request := &iam.CreateServiceAccountRequest{
		ServiceAccount: &iam.ServiceAccount{
			Name:        ServiceAccountName,
			DisplayName: ServiceAccountDisplayName,
			Description: ServiceAccountDescription,
		},
	}

	serviceAccount, err := client.CreateServiceAccount(ctx, request)
	if err != nil {
		log.Fatalf("CreateServiceAccount: %s", err)
	}

	log.Printf("name         : %s", serviceAccount.GetName())
	log.Printf("uid          : %s", serviceAccount.GetUid())
	log.Printf("display name : %s", serviceAccount.GetDisplayName())
	log.Printf("description  : %s", serviceAccount.GetDescription())
	log.Printf("client id    : %s", serviceAccount.GetOauth2ClientId())
	log.Printf("client secret: %s", serviceAccount.GetOauth2ClientSecret())
}
