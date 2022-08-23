package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/identity/v1alpha1"
	identity "github.com/animeapis/go-genproto/identity/v1alpha1"
)

var (
	User = "[USER]"

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

	request := &identity.GetUserRequest{
		Name: User,
	}

	response, err := client.GetUser(ctx, request)
	if err != nil {
		log.Fatalf("GetUser: %s", err)
	}

	log.Printf("user           : %s", User)
	log.Printf("uuid           : %s", response.Uuid)
	log.Printf("username       : %s", response.Username)
	log.Printf("discriminator  : %s", response.Discriminator)
	log.Printf("primary email  : %s", response.PrimaryEmail)
	log.Printf("email verified : %t", response.PrimaryEmailVerified)
	log.Printf("given name     : %s", response.GivenName)
	log.Printf("family name    : %s", response.FamilyName)
	log.Printf("display name   : %s", response.DisplayName)
	log.Printf("country code   : %s", response.CountryCode)
	log.Printf("locale         : %s", response.Locale)
	log.Printf("locale fallback: %s", response.LocaleFallback)
	log.Printf("birthday       : %s", response.Birthday.String())
	log.Printf("gender         : %s", response.Gender.String())
	log.Printf("profile image  : %s", response.ProfileImage)
	log.Printf("banner image   : %s", response.BannerImage)
}
