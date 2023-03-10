package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/batchdeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/batchendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/codecontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/codeversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/componentcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/componentversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/datacontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/datastore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/dataversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/environmentcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/environmentversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/job"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/machinelearningcomputes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/modelcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/modelversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/onlinedeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/onlineendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/operationalizationclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/quota"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/v2workspaceconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/virtualmachinesizes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/workspaceprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/workspaceprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/workspaces"
)

type Client struct {
	BatchDeployment                      *batchdeployment.BatchDeploymentClient
	BatchEndpoint                        *batchendpoint.BatchEndpointClient
	CodeContainer                        *codecontainer.CodeContainerClient
	CodeVersion                          *codeversion.CodeVersionClient
	ComponentContainer                   *componentcontainer.ComponentContainerClient
	ComponentVersion                     *componentversion.ComponentVersionClient
	DataContainer                        *datacontainer.DataContainerClient
	DataVersion                          *dataversion.DataVersionClient
	Datastore                            *datastore.DatastoreClient
	EnvironmentContainer                 *environmentcontainer.EnvironmentContainerClient
	EnvironmentVersion                   *environmentversion.EnvironmentVersionClient
	Job                                  *job.JobClient
	MachineLearningComputes              *machinelearningcomputes.MachineLearningComputesClient
	ModelContainer                       *modelcontainer.ModelContainerClient
	ModelVersion                         *modelversion.ModelVersionClient
	OnlineDeployment                     *onlinedeployment.OnlineDeploymentClient
	OnlineEndpoint                       *onlineendpoint.OnlineEndpointClient
	OperationalizationClusters           *operationalizationclusters.OperationalizationClustersClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	Quota                                *quota.QuotaClient
	V2WorkspaceConnectionResource        *v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient
	VirtualMachineSizes                  *virtualmachinesizes.VirtualMachineSizesClient
	WorkspacePrivateEndpointConnections  *workspaceprivateendpointconnections.WorkspacePrivateEndpointConnectionsClient
	WorkspacePrivateLinkResources        *workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	batchDeploymentClient := batchdeployment.NewBatchDeploymentClientWithBaseURI(endpoint)
	configureAuthFunc(&batchDeploymentClient.Client)

	batchEndpointClient := batchendpoint.NewBatchEndpointClientWithBaseURI(endpoint)
	configureAuthFunc(&batchEndpointClient.Client)

	codeContainerClient := codecontainer.NewCodeContainerClientWithBaseURI(endpoint)
	configureAuthFunc(&codeContainerClient.Client)

	codeVersionClient := codeversion.NewCodeVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&codeVersionClient.Client)

	componentContainerClient := componentcontainer.NewComponentContainerClientWithBaseURI(endpoint)
	configureAuthFunc(&componentContainerClient.Client)

	componentVersionClient := componentversion.NewComponentVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&componentVersionClient.Client)

	dataContainerClient := datacontainer.NewDataContainerClientWithBaseURI(endpoint)
	configureAuthFunc(&dataContainerClient.Client)

	dataVersionClient := dataversion.NewDataVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&dataVersionClient.Client)

	datastoreClient := datastore.NewDatastoreClientWithBaseURI(endpoint)
	configureAuthFunc(&datastoreClient.Client)

	environmentContainerClient := environmentcontainer.NewEnvironmentContainerClientWithBaseURI(endpoint)
	configureAuthFunc(&environmentContainerClient.Client)

	environmentVersionClient := environmentversion.NewEnvironmentVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&environmentVersionClient.Client)

	jobClient := job.NewJobClientWithBaseURI(endpoint)
	configureAuthFunc(&jobClient.Client)

	machineLearningComputesClient := machinelearningcomputes.NewMachineLearningComputesClientWithBaseURI(endpoint)
	configureAuthFunc(&machineLearningComputesClient.Client)

	modelContainerClient := modelcontainer.NewModelContainerClientWithBaseURI(endpoint)
	configureAuthFunc(&modelContainerClient.Client)

	modelVersionClient := modelversion.NewModelVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&modelVersionClient.Client)

	onlineDeploymentClient := onlinedeployment.NewOnlineDeploymentClientWithBaseURI(endpoint)
	configureAuthFunc(&onlineDeploymentClient.Client)

	onlineEndpointClient := onlineendpoint.NewOnlineEndpointClientWithBaseURI(endpoint)
	configureAuthFunc(&onlineEndpointClient.Client)

	operationalizationClustersClient := operationalizationclusters.NewOperationalizationClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&operationalizationClustersClient.Client)

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	quotaClient := quota.NewQuotaClientWithBaseURI(endpoint)
	configureAuthFunc(&quotaClient.Client)

	v2WorkspaceConnectionResourceClient := v2workspaceconnectionresource.NewV2WorkspaceConnectionResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&v2WorkspaceConnectionResourceClient.Client)

	virtualMachineSizesClient := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineSizesClient.Client)

	workspacePrivateEndpointConnectionsClient := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacePrivateEndpointConnectionsClient.Client)

	workspacePrivateLinkResourcesClient := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacePrivateLinkResourcesClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		BatchDeployment:                      &batchDeploymentClient,
		BatchEndpoint:                        &batchEndpointClient,
		CodeContainer:                        &codeContainerClient,
		CodeVersion:                          &codeVersionClient,
		ComponentContainer:                   &componentContainerClient,
		ComponentVersion:                     &componentVersionClient,
		DataContainer:                        &dataContainerClient,
		DataVersion:                          &dataVersionClient,
		Datastore:                            &datastoreClient,
		EnvironmentContainer:                 &environmentContainerClient,
		EnvironmentVersion:                   &environmentVersionClient,
		Job:                                  &jobClient,
		MachineLearningComputes:              &machineLearningComputesClient,
		ModelContainer:                       &modelContainerClient,
		ModelVersion:                         &modelVersionClient,
		OnlineDeployment:                     &onlineDeploymentClient,
		OnlineEndpoint:                       &onlineEndpointClient,
		OperationalizationClusters:           &operationalizationClustersClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		Quota:                                &quotaClient,
		V2WorkspaceConnectionResource:        &v2WorkspaceConnectionResourceClient,
		VirtualMachineSizes:                  &virtualMachineSizesClient,
		WorkspacePrivateEndpointConnections:  &workspacePrivateEndpointConnectionsClient,
		WorkspacePrivateLinkResources:        &workspacePrivateLinkResourcesClient,
		Workspaces:                           &workspacesClient,
	}
}
