package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	gapic "github.com/animeapis/api-go-client/iam/admin/v1alpha1"
	iam "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
)

var (
	ServiceAccountName = "[NAME]"

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

	request := &iam.DeleteServiceAccountRequest{
		Name: ServiceAccountName,
	}

	if err := client.DeleteServiceAccount(ctx, request); err != nil {
		if st, ok := status.FromError(err); ok && st.Code() == codes.NotFound {
			log.Fatalf("service account %s not found: %s", ServiceAccountName, st.Message())
		}
		log.Fatalf("CreateServiceAccount: %s", err)
	}

	log.Printf("service account %s successfully deleted", ServiceAccountName)
}
