package v2021_07_01

import (
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/machinelearningcomputes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/operationalizationclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/quota"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/virtualmachinesizes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaceconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaceprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaceprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2021-07-01/workspaceskus"
)

type Client struct {
	MachineLearningComputes              *machinelearningcomputes.MachineLearningComputesClient
	OperationalizationClusters           *operationalizationclusters.OperationalizationClustersClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	Quota                                *quota.QuotaClient
	VirtualMachineSizes                  *virtualmachinesizes.VirtualMachineSizesClient
	WorkspaceConnections                 *workspaceconnections.WorkspaceConnectionsClient
	WorkspacePrivateEndpointConnections  *workspaceprivateendpointconnections.WorkspacePrivateEndpointConnectionsClient
	WorkspacePrivateLinkResources        *workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient
	WorkspaceSkus                        *workspaceskus.WorkspaceSkusClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	machineLearningComputesClient := machinelearningcomputes.NewMachineLearningComputesClientWithBaseURI(endpoint)
	configureAuthFunc(&machineLearningComputesClient.Client)

	operationalizationClustersClient := operationalizationclusters.NewOperationalizationClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&operationalizationClustersClient.Client)

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	quotaClient := quota.NewQuotaClientWithBaseURI(endpoint)
	configureAuthFunc(&quotaClient.Client)

	virtualMachineSizesClient := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineSizesClient.Client)

	workspaceConnectionsClient := workspaceconnections.NewWorkspaceConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceConnectionsClient.Client)

	workspacePrivateEndpointConnectionsClient := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacePrivateEndpointConnectionsClient.Client)

	workspacePrivateLinkResourcesClient := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacePrivateLinkResourcesClient.Client)

	workspaceSkusClient := workspaceskus.NewWorkspaceSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceSkusClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		MachineLearningComputes:              &machineLearningComputesClient,
		OperationalizationClusters:           &operationalizationClustersClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		Quota:                                &quotaClient,
		VirtualMachineSizes:                  &virtualMachineSizesClient,
		WorkspaceConnections:                 &workspaceConnectionsClient,
		WorkspacePrivateEndpointConnections:  &workspacePrivateEndpointConnectionsClient,
		WorkspacePrivateLinkResources:        &workspacePrivateLinkResourcesClient,
		WorkspaceSkus:                        &workspaceSkusClient,
		Workspaces:                           &workspacesClient,
	}
}
