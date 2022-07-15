package v2020_10_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2020-10-01/clusters"
)

type Client struct {
	Clusters *clusters.ClustersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	return Client{
		Clusters: &clustersClient,
	}
}
