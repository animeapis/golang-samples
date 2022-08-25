package main

import (
	"context"
	"crypto/x509"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	gapic "github.com/animeapis/api-go-client/identity/v1alpha1"
	identity "github.com/animeapis/go-genproto/identity/v1alpha1"
)

var (
	User = "[USER]"
)

var (
	Endpoint = "identity.animeapis.com:443"
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

	client, err := gapic.NewClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewClient: %s", err)
	}

	request := &identity.GetUserProfileRequest{
		Name: User,
	}

	response, err := client.GetUserProfile(ctx, request)
	if err != nil {
		log.Fatalf("GetUserProfile: %s", err)
	}

	log.Printf("user         : %s", User)
	log.Printf("username     : %s", response.Username)
	log.Printf("discriminator: %s", response.Discriminator)
	log.Printf("birthday     : %s", response.Birthday.String())
	log.Printf("gender       : %s", response.Gender.String())
	log.Printf("profile image: %s", response.ProfileImage)
	log.Printf("banner image : %s", response.BannerImage)
}
