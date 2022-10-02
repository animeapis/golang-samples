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

	request := &crossref.CreateCrossRefRequest{
		Crossref: &crossref.CrossRef{
			Name: "animes/123",
			Root: "namespaces/animeshon-com/animes/123",
			Edges: []*crossref.CrossRefEdge{
				{
					Name:  "namespaces/animeshon-com/animes/123",
					State: 2,
				},
				{
					Name:  "namespaces/anidb-net/animes/zxc",
					State: 2,
				},
			},
		},
	}

	crossref, err := client.CreateCrossRef(ctx, request)
	if nil != err {
		log.Fatalf("CreateCrossRef: %s", err)
	}

	log.Println("---------------------------------------------------------")
	log.Printf("name        : %s", crossref.GetName())
	log.Printf("root        : %s", crossref.GetRoot())
	log.Printf("operator    : %s", crossref.GetOperator())

	for i, edge := range crossref.GetEdges() {
		log.Printf("edge[%d]name    : %s", i, edge.GetName())
		log.Printf("edge[%d]state   : %d", i, edge.GetState())
	}
}
