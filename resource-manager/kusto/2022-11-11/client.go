package v2022_11_11

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/attacheddatabaseconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/clusterprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/databaseprincipalassignments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/dataconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/kusto"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/managedprivateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/kusto/2022-11-11/scripts"
)

type Client struct {
	AttachedDatabaseConfigurations       *attacheddatabaseconfigurations.AttachedDatabaseConfigurationsClient
	ClusterPrincipalAssignments          *clusterprincipalassignments.ClusterPrincipalAssignmentsClient
	Clusters                             *clusters.ClustersClient
	DataConnections                      *dataconnections.DataConnectionsClient
	DatabasePrincipalAssignments         *databaseprincipalassignments.DatabasePrincipalAssignmentsClient
	Databases                            *databases.DatabasesClient
	Kusto                                *kusto.KustoClient
	ManagedPrivateEndpoints              *managedprivateendpoints.ManagedPrivateEndpointsClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	Scripts                              *scripts.ScriptsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	attachedDatabaseConfigurationsClient := attacheddatabaseconfigurations.NewAttachedDatabaseConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&attachedDatabaseConfigurationsClient.Client)

	clusterPrincipalAssignmentsClient := clusterprincipalassignments.NewClusterPrincipalAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterPrincipalAssignmentsClient.Client)

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	dataConnectionsClient := dataconnections.NewDataConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataConnectionsClient.Client)

	databasePrincipalAssignmentsClient := databaseprincipalassignments.NewDatabasePrincipalAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&databasePrincipalAssignmentsClient.Client)

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	kustoClient := kusto.NewKustoClientWithBaseURI(endpoint)
	configureAuthFunc(&kustoClient.Client)

	managedPrivateEndpointsClient := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedPrivateEndpointsClient.Client)

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	scriptsClient := scripts.NewScriptsClientWithBaseURI(endpoint)
	configureAuthFunc(&scriptsClient.Client)

	return Client{
		AttachedDatabaseConfigurations:       &attachedDatabaseConfigurationsClient,
		ClusterPrincipalAssignments:          &clusterPrincipalAssignmentsClient,
		Clusters:                             &clustersClient,
		DataConnections:                      &dataConnectionsClient,
		DatabasePrincipalAssignments:         &databasePrincipalAssignmentsClient,
		Databases:                            &databasesClient,
		Kusto:                                &kustoClient,
		ManagedPrivateEndpoints:              &managedPrivateEndpointsClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		PrivateLinkResources:                 &privateLinkResourcesClient,
		Scripts:                              &scriptsClient,
	}
}
