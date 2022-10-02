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

	request := &crossref.GetWormholeRequest{
		Name: WormholeName,
	}

	wormhole, err := client.GetWormhole(ctx, request)
	if nil != err {
		log.Fatalf("GetWormhole: %s", err)
	}

	log.Println("---------------------------------------------------------")
	log.Printf("name               : %s", wormhole.GetName())
	for i, name := range wormhole.GetNames() {
		log.Printf("name[%d]text           : %s", i, name.GetText())
		log.Printf("name[%d]loc           : %s", i, name.GetLocalization())
	}
	log.Printf("imageUrl           : %s", wormhole.GetImageUrl())
	log.Printf("type               : %s", wormhole.GetType())
	log.Printf("subType            : %s", wormhole.GetSubtype())
	log.Printf("externalUrl        : %s", wormhole.GetExternalUrl())
	log.Printf("parentName         : %s", wormhole.GetParentName())
	log.Printf("parentExternalUrl  : %s", wormhole.GetParentExternalUrl())
}
