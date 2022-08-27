package main

import (
	"context"
	"log"

	"golang.org/x/oauth2/clientcredentials"

	"google.golang.org/api/option"

	gapic "github.com/animeapis/api-go-client/vision/v1alpha1"
	vision "github.com/animeapis/go-genproto/vision/v1alpha1"
)

var (
	ImageAnnotationParent = "[USER]"
	Resource              = "[RESOURCE]"

	ClientID     = "[CLIENT-ID]"
	ClientSecret = "[CLIENT-SECRET]"
)

var (
	TokenURL = "https://accounts.animeshon.com/o/oauth2/token"
	Endpoint = "vision.animeapis.com:443"
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

	client, err := gapic.NewImageAnnotatorClient(ctx, options...)
	if err != nil {
		log.Fatalf("NewImageAnnotatorClient: %s", err)
	}

	request := &vision.CreateImageAnnotationRequest{
		Parent: ImageAnnotationParent,
		Annotation: &vision.ImageAnnotation{
			Resource: Resource,
			Annotations: &vision.ImageAnnotations{
				// ---------------------------------------------------------------------
				// Annotations related to text in the image.
				// ---------------------------------------------------------------------
				TextAnnotations: []*vision.TextAnnotation{
					{
						Property: &vision.TextAnnotation_TextProperty{
							Languages: []*vision.TextAnnotation_Language{
								{LanguageCode: "eng", Confidence: 0.9},
							},
						},
						BoundingBox: &vision.BoundingPoly{
							Vertices: []*vision.Vertex{
								{X: 0, Y: 0},
								{X: 0, Y: 10},
								{X: 10, Y: 0},
								{X: 10, Y: 10},
							},
						},
						Text:       "Test",
						Confidence: 0.8,
					},
				},
				// ---------------------------------------------------------------------
				// Annotations related to labels such as actions, sentiment, and so on.
				// ---------------------------------------------------------------------
				LabelAnnotations: []*vision.LabelAnnotation{
					{
						Name:       "nsfw",
						Score:      0.8,
						Topicality: 0.9,
					},
				},
				// ---------------------------------------------------------------------
				// Annotations related to entities that can be located in the image.
				// ---------------------------------------------------------------------
				EntityAnnotations: []*vision.EntityAnnotation{
					{
						Name:  "human-eye",
						Score: 0.9,
						BoundingBox: &vision.BoundingPoly{
							Vertices: []*vision.Vertex{
								{X: 0, Y: 0},
								{X: 0, Y: 10},
								{X: 10, Y: 0},
								{X: 10, Y: 10},
							},
						},
					},
				},
				// ---------------------------------------------------------------------
				// Annotations related to known animes, characters, and so on.
				// ---------------------------------------------------------------------
				KnowledgeGraphAnnotations: []*vision.KnowledgeGraphAnnotation{
					{
						Resource:    "animes/12345",
						Score:       0.9,
						BoundingBox: nil,
					},
				},
				// ---------------------------------------------------------------------
				// Annotations related to known links associated to this image.
				// ---------------------------------------------------------------------
				WebSearchAnnotations: []*vision.WebSearchAnnotation{
					{
						Url:   "https://www.pixiv.net/my-link",
						Score: 1.0,
					},
				},
				// ---------------------------------------------------------------------
				// Annotations related to safe search parameters.
				// ---------------------------------------------------------------------
				SafeSearchAnnotation: &vision.SafeSearchAnnotation{
					Adult:    vision.Likelihood_POSSIBLE,
					Racy:     vision.Likelihood_POSSIBLE,
					Violence: vision.Likelihood_POSSIBLE,
					Medical:  vision.Likelihood_POSSIBLE,
					Juvenile: vision.Likelihood_POSSIBLE,
				},
			},
		},
	}

	response, err := client.CreateImageAnnotation(ctx, request)
	if err != nil {
		log.Fatalf("CreateImageAnnotation: %s", err)
	}

	log.Printf("name                         : %s", response.GetName())
	log.Printf("resource                     : %s", response.GetResource())
	log.Printf("[annotations] TEXT           : %+v", response.GetAnnotations().GetTextAnnotations())
	log.Printf("[annotations] LABEL          : %+v", response.GetAnnotations().GetLabelAnnotations())
	log.Printf("[annotations] ENTITY         : %+v", response.GetAnnotations().GetEntityAnnotations())
	log.Printf("[annotations] KNOWLEDGE GRAPH: %+v", response.GetAnnotations().GetKnowledgeGraphAnnotations())
	log.Printf("[annotations] WEB SEARCH     : %+v", response.GetAnnotations().GetWebSearchAnnotations())
	log.Printf("[annotations] SAFE SEARCH    : %+v", response.GetAnnotations().GetSafeSearchAnnotation())
}
