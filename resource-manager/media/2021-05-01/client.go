package v2021_05_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/media/2021-05-01/media"
)

type Client struct {
	Media *media.MediaClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	mediaClient := media.NewMediaClientWithBaseURI(endpoint)
	configureAuthFunc(&mediaClient.Client)

	return Client{
		Media: &mediaClient,
	}
}
