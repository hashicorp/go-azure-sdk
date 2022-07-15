package v2020_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/liveevents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/liveoutputs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/media"
	"github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/streamingendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/media/2020-05-01/streamingendpoints"
)

type Client struct {
	LiveEvents         *liveevents.LiveEventsClient
	LiveOutputs        *liveoutputs.LiveOutputsClient
	Media              *media.MediaClient
	StreamingEndpoint  *streamingendpoint.StreamingEndpointClient
	StreamingEndpoints *streamingendpoints.StreamingEndpointsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	liveEventsClient := liveevents.NewLiveEventsClientWithBaseURI(endpoint)
	configureAuthFunc(&liveEventsClient.Client)

	liveOutputsClient := liveoutputs.NewLiveOutputsClientWithBaseURI(endpoint)
	configureAuthFunc(&liveOutputsClient.Client)

	mediaClient := media.NewMediaClientWithBaseURI(endpoint)
	configureAuthFunc(&mediaClient.Client)

	streamingEndpointClient := streamingendpoint.NewStreamingEndpointClientWithBaseURI(endpoint)
	configureAuthFunc(&streamingEndpointClient.Client)

	streamingEndpointsClient := streamingendpoints.NewStreamingEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&streamingEndpointsClient.Client)

	return Client{
		LiveEvents:         &liveEventsClient,
		LiveOutputs:        &liveOutputsClient,
		Media:              &mediaClient,
		StreamingEndpoint:  &streamingEndpointClient,
		StreamingEndpoints: &streamingEndpointsClient,
	}
}
