package v2021_01_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/key"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/proxy"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mixedreality/2021-01-01/resource"
)

type Client struct {
	Key      *key.KeyClient
	Proxy    *proxy.ProxyClient
	Resource *resource.ResourceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	keyClient := key.NewKeyClientWithBaseURI(endpoint)
	configureAuthFunc(&keyClient.Client)

	proxyClient := proxy.NewProxyClientWithBaseURI(endpoint)
	configureAuthFunc(&proxyClient.Client)

	resourceClient := resource.NewResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&resourceClient.Client)

	return Client{
		Key:      &keyClient,
		Proxy:    &proxyClient,
		Resource: &resourceClient,
	}
}
