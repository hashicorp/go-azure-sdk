package v2022_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/consumption/2022-06-01/pricesheet20220601"
)

type Client struct {
	PriceSheet20220601 *pricesheet20220601.PriceSheet20220601Client
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	priceSheet20220601Client := pricesheet20220601.NewPriceSheet20220601ClientWithBaseURI(endpoint)
	configureAuthFunc(&priceSheet20220601Client.Client)

	return Client{
		PriceSheet20220601: &priceSheet20220601Client,
	}
}
