package v2022_04_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/redhatopenshift/2022-04-01/openshiftclusters"
)

type Client struct {
	OpenShiftClusters *openshiftclusters.OpenShiftClustersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	openShiftClustersClient := openshiftclusters.NewOpenShiftClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&openShiftClustersClient.Client)

	return Client{
		OpenShiftClusters: &openShiftClustersClient,
	}
}
