package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/timestamppb"

	gapic "github.com/animeapis/api-go-client/tracker/v1alpha1"
	tracker "github.com/animeapis/go-genproto/tracker/v1alpha1"
)

var (
	Tracker  = "[TRACKER]"
	Resource = "[RESOURCE]"

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

	request := &tracker.CreateActivityRequest{
		Parent: Tracker,
		Activity: &tracker.Activity{
			Resource:  Resource,
			Platform:  "NETFLIX",
			From:      0,
			To:        22 * 60,
			StartTime: timestamppb.Now(),
			EndTime:   timestamppb.Now(),
		},
	}

	response, err := client.CreateActivity(ctx, request)
	if err != nil {
		log.Fatalf("CreateActivity: %s", err)
	}

	log.Printf("name      : %s", response.GetName())
	log.Printf("resource  : %s", response.GetResource())
	log.Printf("platform  : %s", response.GetPlatform())
	log.Printf("from      : %d", response.GetFrom())
	log.Printf("to        : %d", response.GetTo())
	log.Printf("start time: %s", response.GetStartTime().AsTime().String())
	log.Printf("end time  : %s", response.GetEndTime().AsTime().String())
}
