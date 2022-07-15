package v2021_05_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoanalyzer/2021-05-01-preview/videoanalyzer"
)

type Client struct {
	VideoAnalyzer *videoanalyzer.VideoAnalyzerClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	videoAnalyzerClient := videoanalyzer.NewVideoAnalyzerClientWithBaseURI(endpoint)
	configureAuthFunc(&videoAnalyzerClient.Client)

	return Client{
		VideoAnalyzer: &videoAnalyzerClient,
	}
}
