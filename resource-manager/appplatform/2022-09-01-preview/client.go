package v2022_09_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/appplatform/2022-09-01-preview/appplatform"
)

type Client struct {
	AppPlatform *appplatform.AppPlatformClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	appPlatformClient := appplatform.NewAppPlatformClientWithBaseURI(endpoint)
	configureAuthFunc(&appPlatformClient.Client)

	return Client{
		AppPlatform: &appPlatformClient,
	}
}
