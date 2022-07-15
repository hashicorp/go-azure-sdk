package v2022_02_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2022-02-01/signalr"
)

type Client struct {
	SignalR *signalr.SignalRClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	signalRClient := signalr.NewSignalRClientWithBaseURI(endpoint)
	configureAuthFunc(&signalRClient.Client)

	return Client{
		SignalR: &signalRClient,
	}
}
