package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/tracker/v1alpha1"
	tracker "github.com/animeapis/go-genproto/tracker/v1alpha1"
)

var (
	TrackerParent = "[TRACKER]"

	Provider = tracker.Provider_MYANIMELIST

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "tracker.animeapis.com:443"
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

	request := &tracker.ImportTrackersRequest{
		Parent:   TrackerParent,
		Provider: Provider,
	}

	response, err := client.ImportTrackers(ctx, request)
	if err != nil {
		log.Fatalf("ImportTrackers: %s", err)
	}

	log.Printf("waiting for import to be completed...")
	if _, err := response.Wait(ctx); err != nil {
		log.Fatalf("Wait: %s", err)
	}

	log.Printf("import successfully completed")
}
