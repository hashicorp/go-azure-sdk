package v2024_10_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/batchdeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/batchendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/capabilityhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/codecontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/codeversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/componentcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/componentversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/datacontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/datacontainerregistry"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/datareference"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/datastore"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/dataversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/dataversionregistry"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/endpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/environmentcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/environmentversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/feature"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/featuresetcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/featuresetversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/featurestoreentitycontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/featurestoreentityversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencedeltamodel"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferenceendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencegroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/inferencepool"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/job"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/machinelearningcomputes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/managednetwork"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/marketplacesubscription"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/modelcontainer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/modelversion"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/onlinedeployment"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/onlineendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/operationalizationclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/proxyoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/quota"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/registrymanagement"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/serverlessendpoint"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/v2workspaceconnectionresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/virtualmachinesizes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/workspaceprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/workspaceprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/machinelearningservices/2024-10-01-preview/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BatchDeployment                      *batchdeployment.BatchDeploymentClient
	BatchEndpoint                        *batchendpoint.BatchEndpointClient
	CapabilityHost                       *capabilityhost.CapabilityHostClient
	CodeContainer                        *codecontainer.CodeContainerClient
	CodeVersion                          *codeversion.CodeVersionClient
	ComponentContainer                   *componentcontainer.ComponentContainerClient
	ComponentVersion                     *componentversion.ComponentVersionClient
	DataContainer                        *datacontainer.DataContainerClient
	DataContainerRegistry                *datacontainerregistry.DataContainerRegistryClient
	DataReference                        *datareference.DataReferenceClient
	DataVersion                          *dataversion.DataVersionClient
	DataVersionRegistry                  *dataversionregistry.DataVersionRegistryClient
	Datastore                            *datastore.DatastoreClient
	Endpoint                             *endpoint.EndpointClient
	EnvironmentContainer                 *environmentcontainer.EnvironmentContainerClient
	EnvironmentVersion                   *environmentversion.EnvironmentVersionClient
	Feature                              *feature.FeatureClient
	FeaturesetContainer                  *featuresetcontainer.FeaturesetContainerClient
	FeaturesetVersion                    *featuresetversion.FeaturesetVersionClient
	FeaturestoreEntityContainer          *featurestoreentitycontainer.FeaturestoreEntityContainerClient
	FeaturestoreEntityVersion            *featurestoreentityversion.FeaturestoreEntityVersionClient
	InferenceDeltaModel                  *inferencedeltamodel.InferenceDeltaModelClient
	InferenceEndpoint                    *inferenceendpoint.InferenceEndpointClient
	InferenceGroup                       *inferencegroup.InferenceGroupClient
	InferencePool                        *inferencepool.InferencePoolClient
	Job                                  *job.JobClient
	MachineLearningComputes              *machinelearningcomputes.MachineLearningComputesClient
	ManagedNetwork                       *managednetwork.ManagedNetworkClient
	MarketplaceSubscription              *marketplacesubscription.MarketplaceSubscriptionClient
	ModelContainer                       *modelcontainer.ModelContainerClient
	ModelVersion                         *modelversion.ModelVersionClient
	OnlineDeployment                     *onlinedeployment.OnlineDeploymentClient
	OnlineEndpoint                       *onlineendpoint.OnlineEndpointClient
	OperationalizationClusters           *operationalizationclusters.OperationalizationClustersClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	ProxyOperations                      *proxyoperations.ProxyOperationsClient
	Quota                                *quota.QuotaClient
	RegistryManagement                   *registrymanagement.RegistryManagementClient
	Schedule                             *schedule.ScheduleClient
	ServerlessEndpoint                   *serverlessendpoint.ServerlessEndpointClient
	V2WorkspaceConnectionResource        *v2workspaceconnectionresource.V2WorkspaceConnectionResourceClient
	VirtualMachineSizes                  *virtualmachinesizes.VirtualMachineSizesClient
	WorkspacePrivateEndpointConnections  *workspaceprivateendpointconnections.WorkspacePrivateEndpointConnectionsClient
	WorkspacePrivateLinkResources        *workspaceprivatelinkresources.WorkspacePrivateLinkResourcesClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	batchDeploymentClient, err := batchdeployment.NewBatchDeploymentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BatchDeployment client: %+v", err)
	}
	configureFunc(batchDeploymentClient.Client)

	batchEndpointClient, err := batchendpoint.NewBatchEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BatchEndpoint client: %+v", err)
	}
	configureFunc(batchEndpointClient.Client)

	capabilityHostClient, err := capabilityhost.NewCapabilityHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapabilityHost client: %+v", err)
	}
	configureFunc(capabilityHostClient.Client)

	codeContainerClient, err := codecontainer.NewCodeContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CodeContainer client: %+v", err)
	}
	configureFunc(codeContainerClient.Client)

	codeVersionClient, err := codeversion.NewCodeVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CodeVersion client: %+v", err)
	}
	configureFunc(codeVersionClient.Client)

	componentContainerClient, err := componentcontainer.NewComponentContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentContainer client: %+v", err)
	}
	configureFunc(componentContainerClient.Client)

	componentVersionClient, err := componentversion.NewComponentVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComponentVersion client: %+v", err)
	}
	configureFunc(componentVersionClient.Client)

	dataContainerClient, err := datacontainer.NewDataContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataContainer client: %+v", err)
	}
	configureFunc(dataContainerClient.Client)

	dataContainerRegistryClient, err := datacontainerregistry.NewDataContainerRegistryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataContainerRegistry client: %+v", err)
	}
	configureFunc(dataContainerRegistryClient.Client)

	dataReferenceClient, err := datareference.NewDataReferenceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataReference client: %+v", err)
	}
	configureFunc(dataReferenceClient.Client)

	dataVersionClient, err := dataversion.NewDataVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataVersion client: %+v", err)
	}
	configureFunc(dataVersionClient.Client)

	dataVersionRegistryClient, err := dataversionregistry.NewDataVersionRegistryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataVersionRegistry client: %+v", err)
	}
	configureFunc(dataVersionRegistryClient.Client)

	datastoreClient, err := datastore.NewDatastoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Datastore client: %+v", err)
	}
	configureFunc(datastoreClient.Client)

	endpointClient, err := endpoint.NewEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Endpoint client: %+v", err)
	}
	configureFunc(endpointClient.Client)

	environmentContainerClient, err := environmentcontainer.NewEnvironmentContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnvironmentContainer client: %+v", err)
	}
	configureFunc(environmentContainerClient.Client)

	environmentVersionClient, err := environmentversion.NewEnvironmentVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EnvironmentVersion client: %+v", err)
	}
	configureFunc(environmentVersionClient.Client)

	featureClient, err := feature.NewFeatureClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Feature client: %+v", err)
	}
	configureFunc(featureClient.Client)

	featuresetContainerClient, err := featuresetcontainer.NewFeaturesetContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeaturesetContainer client: %+v", err)
	}
	configureFunc(featuresetContainerClient.Client)

	featuresetVersionClient, err := featuresetversion.NewFeaturesetVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeaturesetVersion client: %+v", err)
	}
	configureFunc(featuresetVersionClient.Client)

	featurestoreEntityContainerClient, err := featurestoreentitycontainer.NewFeaturestoreEntityContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeaturestoreEntityContainer client: %+v", err)
	}
	configureFunc(featurestoreEntityContainerClient.Client)

	featurestoreEntityVersionClient, err := featurestoreentityversion.NewFeaturestoreEntityVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FeaturestoreEntityVersion client: %+v", err)
	}
	configureFunc(featurestoreEntityVersionClient.Client)

	inferenceDeltaModelClient, err := inferencedeltamodel.NewInferenceDeltaModelClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InferenceDeltaModel client: %+v", err)
	}
	configureFunc(inferenceDeltaModelClient.Client)

	inferenceEndpointClient, err := inferenceendpoint.NewInferenceEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InferenceEndpoint client: %+v", err)
	}
	configureFunc(inferenceEndpointClient.Client)

	inferenceGroupClient, err := inferencegroup.NewInferenceGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InferenceGroup client: %+v", err)
	}
	configureFunc(inferenceGroupClient.Client)

	inferencePoolClient, err := inferencepool.NewInferencePoolClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InferencePool client: %+v", err)
	}
	configureFunc(inferencePoolClient.Client)

	jobClient, err := job.NewJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Job client: %+v", err)
	}
	configureFunc(jobClient.Client)

	machineLearningComputesClient, err := machinelearningcomputes.NewMachineLearningComputesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineLearningComputes client: %+v", err)
	}
	configureFunc(machineLearningComputesClient.Client)

	managedNetworkClient, err := managednetwork.NewManagedNetworkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedNetwork client: %+v", err)
	}
	configureFunc(managedNetworkClient.Client)

	marketplaceSubscriptionClient, err := marketplacesubscription.NewMarketplaceSubscriptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MarketplaceSubscription client: %+v", err)
	}
	configureFunc(marketplaceSubscriptionClient.Client)

	modelContainerClient, err := modelcontainer.NewModelContainerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ModelContainer client: %+v", err)
	}
	configureFunc(modelContainerClient.Client)

	modelVersionClient, err := modelversion.NewModelVersionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ModelVersion client: %+v", err)
	}
	configureFunc(modelVersionClient.Client)

	onlineDeploymentClient, err := onlinedeployment.NewOnlineDeploymentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineDeployment client: %+v", err)
	}
	configureFunc(onlineDeploymentClient.Client)

	onlineEndpointClient, err := onlineendpoint.NewOnlineEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OnlineEndpoint client: %+v", err)
	}
	configureFunc(onlineEndpointClient.Client)

	operationalizationClustersClient, err := operationalizationclusters.NewOperationalizationClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OperationalizationClusters client: %+v", err)
	}
	configureFunc(operationalizationClustersClient.Client)

	outboundNetworkDependenciesEndpointsClient, err := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundNetworkDependenciesEndpoints client: %+v", err)
	}
	configureFunc(outboundNetworkDependenciesEndpointsClient.Client)

	proxyOperationsClient, err := proxyoperations.NewProxyOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProxyOperations client: %+v", err)
	}
	configureFunc(proxyOperationsClient.Client)

	quotaClient, err := quota.NewQuotaClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Quota client: %+v", err)
	}
	configureFunc(quotaClient.Client)

	registryManagementClient, err := registrymanagement.NewRegistryManagementClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RegistryManagement client: %+v", err)
	}
	configureFunc(registryManagementClient.Client)

	scheduleClient, err := schedule.NewScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %+v", err)
	}
	configureFunc(scheduleClient.Client)

	serverlessEndpointClient, err := serverlessendpoint.NewServerlessEndpointClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerlessEndpoint client: %+v", err)
	}
	configureFunc(serverlessEndpointClient.Client)

	v2WorkspaceConnectionResourceClient, err := v2workspaceconnectionresource.NewV2WorkspaceConnectionResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building V2WorkspaceConnectionResource client: %+v", err)
	}
	configureFunc(v2WorkspaceConnectionResourceClient.Client)

	virtualMachineSizesClient, err := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineSizes client: %+v", err)
	}
	configureFunc(virtualMachineSizesClient.Client)

	workspacePrivateEndpointConnectionsClient, err := workspaceprivateendpointconnections.NewWorkspacePrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspacePrivateEndpointConnections client: %+v", err)
	}
	configureFunc(workspacePrivateEndpointConnectionsClient.Client)

	workspacePrivateLinkResourcesClient, err := workspaceprivatelinkresources.NewWorkspacePrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspacePrivateLinkResources client: %+v", err)
	}
	configureFunc(workspacePrivateLinkResourcesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		BatchDeployment:                      batchDeploymentClient,
		BatchEndpoint:                        batchEndpointClient,
		CapabilityHost:                       capabilityHostClient,
		CodeContainer:                        codeContainerClient,
		CodeVersion:                          codeVersionClient,
		ComponentContainer:                   componentContainerClient,
		ComponentVersion:                     componentVersionClient,
		DataContainer:                        dataContainerClient,
		DataContainerRegistry:                dataContainerRegistryClient,
		DataReference:                        dataReferenceClient,
		DataVersion:                          dataVersionClient,
		DataVersionRegistry:                  dataVersionRegistryClient,
		Datastore:                            datastoreClient,
		Endpoint:                             endpointClient,
		EnvironmentContainer:                 environmentContainerClient,
		EnvironmentVersion:                   environmentVersionClient,
		Feature:                              featureClient,
		FeaturesetContainer:                  featuresetContainerClient,
		FeaturesetVersion:                    featuresetVersionClient,
		FeaturestoreEntityContainer:          featurestoreEntityContainerClient,
		FeaturestoreEntityVersion:            featurestoreEntityVersionClient,
		InferenceDeltaModel:                  inferenceDeltaModelClient,
		InferenceEndpoint:                    inferenceEndpointClient,
		InferenceGroup:                       inferenceGroupClient,
		InferencePool:                        inferencePoolClient,
		Job:                                  jobClient,
		MachineLearningComputes:              machineLearningComputesClient,
		ManagedNetwork:                       managedNetworkClient,
		MarketplaceSubscription:              marketplaceSubscriptionClient,
		ModelContainer:                       modelContainerClient,
		ModelVersion:                         modelVersionClient,
		OnlineDeployment:                     onlineDeploymentClient,
		OnlineEndpoint:                       onlineEndpointClient,
		OperationalizationClusters:           operationalizationClustersClient,
		OutboundNetworkDependenciesEndpoints: outboundNetworkDependenciesEndpointsClient,
		ProxyOperations:                      proxyOperationsClient,
		Quota:                                quotaClient,
		RegistryManagement:                   registryManagementClient,
		Schedule:                             scheduleClient,
		ServerlessEndpoint:                   serverlessEndpointClient,
		V2WorkspaceConnectionResource:        v2WorkspaceConnectionResourceClient,
		VirtualMachineSizes:                  virtualMachineSizesClient,
		WorkspacePrivateEndpointConnections:  workspacePrivateEndpointConnectionsClient,
		WorkspacePrivateLinkResources:        workspacePrivateLinkResourcesClient,
		Workspaces:                           workspacesClient,
	}, nil
}
