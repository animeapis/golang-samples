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
	TrackerParent = "[USER-OR-AUDIENCE]"
	Resource      = "[RESOURCE]"

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

	request := &tracker.CreateTrackerRequest{
		Parent: TrackerParent,
		Tracker: &tracker.Tracker{
			Resource:  Resource,
			StartTime: timestamppb.Now(),
			EndTime:   timestamppb.Now(),
			State:     tracker.State_IN_PROGRESS,
		},
	}

	response, err := client.CreateTracker(ctx, request)
	if err != nil {
		log.Fatalf("CreateTracker: %s", err)
	}

	log.Printf("name      : %s", response.GetName())
	log.Printf("resource  : %s", response.GetResource())
	log.Printf("completed : %v", response.GetCompletedResources())
	log.Printf("start time: %s", response.GetStartTime().AsTime().String())
	log.Printf("end time  : %s", response.GetEndTime().AsTime().String())
	log.Printf("progress  : %.2f", response.GetProgressPercentage().GetValue())
	log.Printf("state     : %s", response.GetState().String())
}
