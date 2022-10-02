package main

import (
	"context"
	"errors"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/crossrefs/v1alpha1"
	crossref "github.com/animeapis/go-genproto/crossrefs/v1alpha1"
)

var (
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

	request := &crossref.ListCrossRefsRequest{
		Filter: &crossref.CrossRefsFilterRequest{
			Prefix: "animes",
		},
		PageSize: 10,
	}

	it := client.ListCrossRefs(ctx, request)
	if nil != err {
		log.Fatalf("ListCrossRefs: %s", err)
	}

	for {
		crossref, err := it.Next()
		if err != nil {
			if errors.Is(err, iterator.Done) {
				log.Println("---------------------------------------------------------")
				return
			}

			log.Fatalf("ListCrossRefs: %s", err)
		}

		log.Printf("name        : %s", crossref.GetName())
		log.Printf("root        : %s", crossref.GetRoot())
		log.Printf("operator    : %s", crossref.GetOperator())

		for i, edge := range crossref.GetEdges() {
			log.Printf("edge[%d]name    : %s", i, edge.GetName())
			log.Printf("edge[%d]state   : %d", i, edge.GetState())
		}
	}
}
