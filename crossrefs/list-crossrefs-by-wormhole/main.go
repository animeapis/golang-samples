package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/crossrefs/v1alpha1"
	crossref "github.com/animeapis/go-genproto/crossrefs/v1alpha1"
)

var (
	WormholeName = "[RESOURCE NAME]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "crossrefs.animeapis.com:443"
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

	client, err := gapic.NewReferrerClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewClient: %s", err)
	}

	request := &crossref.ListWormholeCrossRefsRequest{
		Name:         WormholeName,
		WithApproved: true,
		WithPending:  true,
		WithPartial:  true,
		WithRejected: true,
	}

	wormholes, err := client.ListWormholeCrossRefs(ctx, request)
	if nil != err {
		log.Fatalf("ListWormholeCrossRefs: %s", err)
	}

	for _, crossref := range wormholes.GetCrossrefs() {
		log.Printf("name               : %s", crossref.GetName())
	}

}
