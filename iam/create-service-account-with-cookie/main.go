package main

import (
	"context"
	"crypto/x509"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"

	gapic "github.com/animeapis/api-go-client/iam/admin/v1alpha1"
	iam "github.com/animeapis/go-genproto/iam/admin/v1alpha1"
)

var (
	ServiceAccountName        = "[NAME]"
	ServiceAccountDisplayName = "[DISPLAY-NAME]"
	ServiceAccountDescription = "[DESCRIPTION]"

	Cookie = "[COOKIE]"
)

var (
	Endpoint = "iam.animeapis.com:443"
)

func main() {
	ctx := context.Background()

	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatalf("SystemCertPool: %s", err)
	}

	transportCredentials := grpc.WithTransportCredentials(
		credentials.NewClientTLSFromCert(pool, ""),
	)

	options := []option.ClientOption{
		option.WithEndpoint(Endpoint),
		option.WithoutAuthentication(),
		option.WithGRPCDialOption(transportCredentials),
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

	// Set the cookie content as `Cookie` header.
	ctx = metadata.AppendToOutgoingContext(ctx, "cookie", Cookie)

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
