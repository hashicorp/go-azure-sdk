package v2021_10_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridkubernetes/2021-10-01/hybridkubernetes"
)

type Client struct {
	HybridKubernetes *hybridkubernetes.HybridKubernetesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	hybridKubernetesClient := hybridkubernetes.NewHybridKubernetesClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridKubernetesClient.Client)

	return Client{
		HybridKubernetes: &hybridKubernetesClient,
	}
}
