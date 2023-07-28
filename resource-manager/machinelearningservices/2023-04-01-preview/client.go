package v2023_04_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/batchdeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/batchendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/codecontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/codeversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/componentcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/componentversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/datacontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/datacontainerregistry"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/datastore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/dataversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/dataversionregistry"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/environmentcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/environmentversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/feature"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featuresetcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featuresetversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featurestoreentitycontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/featurestoreentityversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/job"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/labelingjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/machinelearningcomputes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/managednetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/modelcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/modelversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/onlinedeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/onlineendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/operationalizationclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/proxyoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/quota"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/registrymanagement"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/v2workspaceconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/virtualmachinesizes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/workspaceprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/workspaceprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2023-04-01-preview/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BatchDeployment                      *batchdeployment.BatchDeploymentClient
	BatchEndpoint                        *batchendpoint.BatchEndpointClient
	CodeContainer                        *codecontainer.CodeContainerClient
	CodeVersion                          *codeversion.CodeVersionClient
	ComponentContainer                   *componentcontainer.ComponentContainerClient
	ComponentVersion                     *componentversion.ComponentVersionClient
	DataContainer                        *datacontainer.DataContainerClient
	DataContainerRegistry                *datacontainerregistry.DataContainerRegistryClient
	DataVersion                          *dataversion.DataVersionClient
	DataVersionRegistry                  *dataversionregistry.DataVersionRegistryClient
	Datastore                            *datastore.DatastoreClient
	EnvironmentContainer                 *environmentcontainer.EnvironmentContainerClient
	EnvironmentVersion                   *environmentversion.EnvironmentVersionClient
	Feature                              *feature.FeatureClient
	FeaturesetContainer                  *featuresetcontainer.FeaturesetContainerClient
	FeaturesetVersion                    *featuresetversion.FeaturesetVersionClient
	FeaturestoreEntityContainer          *featurestoreentitycontainer.FeaturestoreEntityContainerClient
	FeaturestoreEntityVersion            *featurestoreentityversion.FeaturestoreEntityVersionClient
	Job                                  *job.JobClient
	LabelingJob                          *labelingjob.LabelingJobClient
	MachineLearningComputes              *machinelearningcomputes.MachineLearningComputesClient
	ManagedNetwork                       *managednetwork.ManagedNetworkClient
	ModelContainer                       *modelcontainer.ModelContainerClient
	ModelVersion                         *modelversion.ModelVersionClient
	OnlineDeployment                     *onlinedeployment.OnlineDeploymentClient
	OnlineEndpoint                       *onlineendpoint.OnlineEndpointClient
	OperationalizationClusters           *operationalizationclusters.OperationalizationClustersClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	ProxyOperations                      *proxyoperations.ProxyOperationsClient
	Quota                                *quota.QuotaClient
	RegistryManagement                   *registrymanagement.RegistryManagementClient
	Schedule                             *schedule.ScheduleClient
	V2WorkspaceConnectionResource        *v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient
	VirtualMachineSizes                  *virtualmachinesizes.VirtualMachineSizesClient
	WorkspacePrivateEndpointConnections  *workspaceprivateendpointconnections.WorkspacePrivateEndpointConnectionsClient
	WorkspacePrivateLinkResources        *workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	batchDeploymentClient, err := batchdeployment.NewBatchDeploymentClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BatchDeployment client: %+v", err)
	}
	configureFunc(batchDeploymentClient.Client)

	batchEndpointClient, err := batchendpoint.NewBatchEndpointClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BatchEndpoint client: %+v", err)
	}
	configureFunc(batchEndpointClient.Client)

	codeContainerClient, err := codecontainer.NewCodeContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CodeContainer client: %+v", err)
	}
	configureFunc(codeContainerClient.Client)

	codeVersionClient, err := codeversion.NewCodeVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CodeVersion client: %+v", err)
	}
	configureFunc(codeVersionClient.Client)

	componentContainerClient, err := componentcontainer.NewComponentContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ComponentContainer client: %+v", err)
	}
	configureFunc(componentContainerClient.Client)

	componentVersionClient, err := componentversion.NewComponentVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ComponentVersion client: %+v", err)
	}
	configureFunc(componentVersionClient.Client)

	dataContainerClient, err := datacontainer.NewDataContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataContainer client: %+v", err)
	}
	configureFunc(dataContainerClient.Client)

	dataContainerRegistryClient, err := datacontainerregistry.NewDataContainerRegistryClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataContainerRegistry client: %+v", err)
	}
	configureFunc(dataContainerRegistryClient.Client)

	dataVersionClient, err := dataversion.NewDataVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataVersion client: %+v", err)
	}
	configureFunc(dataVersionClient.Client)

	dataVersionRegistryClient, err := dataversionregistry.NewDataVersionRegistryClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataVersionRegistry client: %+v", err)
	}
	configureFunc(dataVersionRegistryClient.Client)

	datastoreClient, err := datastore.NewDatastoreClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Datastore client: %+v", err)
	}
	configureFunc(datastoreClient.Client)

	environmentContainerClient, err := environmentcontainer.NewEnvironmentContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EnvironmentContainer client: %+v", err)
	}
	configureFunc(environmentContainerClient.Client)

	environmentVersionClient, err := environmentversion.NewEnvironmentVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building EnvironmentVersion client: %+v", err)
	}
	configureFunc(environmentVersionClient.Client)

	featureClient, err := feature.NewFeatureClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Feature client: %+v", err)
	}
	configureFunc(featureClient.Client)

	featuresetContainerClient, err := featuresetcontainer.NewFeaturesetContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FeaturesetContainer client: %+v", err)
	}
	configureFunc(featuresetContainerClient.Client)

	featuresetVersionClient, err := featuresetversion.NewFeaturesetVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FeaturesetVersion client: %+v", err)
	}
	configureFunc(featuresetVersionClient.Client)

	featurestoreEntityContainerClient, err := featurestoreentitycontainer.NewFeaturestoreEntityContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FeaturestoreEntityContainer client: %+v", err)
	}
	configureFunc(featurestoreEntityContainerClient.Client)

	featurestoreEntityVersionClient, err := featurestoreentityversion.NewFeaturestoreEntityVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FeaturestoreEntityVersion client: %+v", err)
	}
	configureFunc(featurestoreEntityVersionClient.Client)

	jobClient, err := job.NewJobClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Job client: %+v", err)
	}
	configureFunc(jobClient.Client)

	labelingJobClient, err := labelingjob.NewLabelingJobClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LabelingJob client: %+v", err)
	}
	configureFunc(labelingJobClient.Client)

	machineLearningComputesClient, err := machinelearningcomputes.NewMachineLearningComputesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building MachineLearningComputes client: %+v", err)
	}
	configureFunc(machineLearningComputesClient.Client)

	managedNetworkClient, err := managednetwork.NewManagedNetworkClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedNetwork client: %+v", err)
	}
	configureFunc(managedNetworkClient.Client)

	modelContainerClient, err := modelcontainer.NewModelContainerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ModelContainer client: %+v", err)
	}
	configureFunc(modelContainerClient.Client)

	modelVersionClient, err := modelversion.NewModelVersionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ModelVersion client: %+v", err)
	}
	configureFunc(modelVersionClient.Client)

	onlineDeploymentClient, err := onlinedeployment.NewOnlineDeploymentClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OnlineDeployment client: %+v", err)
	}
	configureFunc(onlineDeploymentClient.Client)

	onlineEndpointClient, err := onlineendpoint.NewOnlineEndpointClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OnlineEndpoint client: %+v", err)
	}
	configureFunc(onlineEndpointClient.Client)

	operationalizationClustersClient, err := operationalizationclusters.NewOperationalizationClustersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OperationalizationClusters client: %+v", err)
	}
	configureFunc(operationalizationClustersClient.Client)

	outboundNetworkDependenciesEndpointsClient, err := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OutboundNetworkDependenciesEndpoints client: %+v", err)
	}
	configureFunc(outboundNetworkDependenciesEndpointsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	proxyOperationsClient, err := proxyoperations.NewProxyOperationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ProxyOperations client: %+v", err)
	}
	configureFunc(proxyOperationsClient.Client)

	quotaClient, err := quota.NewQuotaClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Quota client: %+v", err)
	}
	configureFunc(quotaClient.Client)

	registryManagementClient, err := registrymanagement.NewRegistryManagementClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RegistryManagement client: %+v", err)
	}
	configureFunc(registryManagementClient.Client)

	scheduleClient, err := schedule.NewScheduleClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %+v", err)
	}
	configureFunc(scheduleClient.Client)

	v2WorkspaceConnectionResourceClient, err := v2workspaceconnectionresource.NewV2WorkspaceConnectionResourceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building V2WorkspaceConnectionResource client: %+v", err)
	}
	configureFunc(v2WorkspaceConnectionResourceClient.Client)

	virtualMachineSizesClient, err := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineSizes client: %+v", err)
	}
	configureFunc(virtualMachineSizesClient.Client)

	workspacePrivateEndpointConnectionsClient, err := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WorkspacePrivateEndpointConnections client: %+v", err)
	}
	configureFunc(workspacePrivateEndpointConnectionsClient.Client)

	workspacePrivateLinkResourcesClient, err := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WorkspacePrivateLinkResources client: %+v", err)
	}
	configureFunc(workspacePrivateLinkResourcesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		BatchDeployment:                      batchDeploymentClient,
		BatchEndpoint:                        batchEndpointClient,
		CodeContainer:                        codeContainerClient,
		CodeVersion:                          codeVersionClient,
		ComponentContainer:                   componentContainerClient,
		ComponentVersion:                     componentVersionClient,
		DataContainer:                        dataContainerClient,
		DataContainerRegistry:                dataContainerRegistryClient,
		DataVersion:                          dataVersionClient,
		DataVersionRegistry:                  dataVersionRegistryClient,
		Datastore:                            datastoreClient,
		EnvironmentContainer:                 environmentContainerClient,
		EnvironmentVersion:                   environmentVersionClient,
		Feature:                              featureClient,
		FeaturesetContainer:                  featuresetContainerClient,
		FeaturesetVersion:                    featuresetVersionClient,
		FeaturestoreEntityContainer:          featurestoreEntityContainerClient,
		FeaturestoreEntityVersion:            featurestoreEntityVersionClient,
		Job:                                  jobClient,
		LabelingJob:                          labelingJobClient,
		MachineLearningComputes:              machineLearningComputesClient,
		ManagedNetwork:                       managedNetworkClient,
		ModelContainer:                       modelContainerClient,
		ModelVersion:                         modelVersionClient,
		OnlineDeployment:                     onlineDeploymentClient,
		OnlineEndpoint:                       onlineEndpointClient,
		OperationalizationClusters:           operationalizationClustersClient,
		OutboundNetworkDependenciesEndpoints: outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           privateEndpointConnectionsClient,
		ProxyOperations:                      proxyOperationsClient,
		Quota:                                quotaClient,
		RegistryManagement:                   registryManagementClient,
		Schedule:                             scheduleClient,
		V2WorkspaceConnectionResource:        v2WorkspaceConnectionResourceClient,
		VirtualMachineSizes:                  virtualMachineSizesClient,
		WorkspacePrivateEndpointConnections:  workspacePrivateEndpointConnectionsClient,
		WorkspacePrivateLinkResources:        workspacePrivateLinkResourcesClient,
		Workspaces:                           workspacesClient,
	}, nil
}
