package v2022_09_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2022-09-01/arcsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2022-09-01/cluster"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2022-09-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2022-09-01/extensions"
)

type Client struct {
	ArcSettings *arcsettings.ArcSettingsClient
	Cluster     *cluster.ClusterClient
	Clusters    *clusters.ClustersClient
	Extensions  *extensions.ExtensionsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	arcSettingsClient := arcsettings.NewArcSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&arcSettingsClient.Client)

	clusterClient := cluster.NewClusterClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterClient.Client)

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	extensionsClient := extensions.NewExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&extensionsClient.Client)

	return Client{
		ArcSettings: &arcSettingsClient,
		Cluster:     &clusterClient,
		Clusters:    &clustersClient,
		Extensions:  &extensionsClient,
	}
}
