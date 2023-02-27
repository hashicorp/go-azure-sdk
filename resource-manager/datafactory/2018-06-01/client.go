package v2018_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/activityruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/credentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/dataflowdebugsession"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/dataflows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/datasets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/exposurecontrol"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/factories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/globalparameters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimenodes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimeobjectmetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/linkedservices"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/managedprivateendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/managedvirtualnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/pipelineruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/pipelines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/trigger"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/triggerruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/triggers"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Activityruns                     *activityruns.ActivityrunsClient
	Credentials                      *credentials.CredentialsClient
	DataFlowDebugSession             *dataflowdebugsession.DataFlowDebugSessionClient
	DataFlows                        *dataflows.DataFlowsClient
	Datasets                         *datasets.DatasetsClient
	ExposureControl                  *exposurecontrol.ExposureControlClient
	Factories                        *factories.FactoriesClient
	GlobalParameters                 *globalparameters.GlobalParametersClient
	IntegrationRuntimeNodes          *integrationruntimenodes.IntegrationRuntimeNodesClient
	IntegrationRuntimeObjectMetadata *integrationruntimeobjectmetadata.IntegrationRuntimeObjectMetadataClient
	IntegrationRuntimes              *integrationruntimes.IntegrationRuntimesClient
	LinkedServices                   *linkedservices.LinkedServicesClient
	ManagedPrivateEndpoints          *managedprivateendpoints.ManagedPrivateEndpointsClient
	ManagedVirtualNetworks           *managedvirtualnetworks.ManagedVirtualNetworksClient
	PipelineRuns                     *pipelineruns.PipelineRunsClient
	Pipelines                        *pipelines.PipelinesClient
	PrivateEndpointConnections       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources             *privatelinkresources.PrivateLinkResourcesClient
	Trigger                          *trigger.TriggerClient
	Triggerruns                      *triggerruns.TriggerrunsClient
	Triggers                         *triggers.TriggersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	activityrunsClient := activityruns.NewActivityrunsClientWithBaseURI(endpoint)
	configureAuthFunc(&activityrunsClient.Client)

	credentialsClient := credentials.NewCredentialsClientWithBaseURI(endpoint)
	configureAuthFunc(&credentialsClient.Client)

	dataFlowDebugSessionClient := dataflowdebugsession.NewDataFlowDebugSessionClientWithBaseURI(endpoint)
	configureAuthFunc(&dataFlowDebugSessionClient.Client)

	dataFlowsClient := dataflows.NewDataFlowsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataFlowsClient.Client)

	datasetsClient := datasets.NewDatasetsClientWithBaseURI(endpoint)
	configureAuthFunc(&datasetsClient.Client)

	exposureControlClient := exposurecontrol.NewExposureControlClientWithBaseURI(endpoint)
	configureAuthFunc(&exposureControlClient.Client)

	factoriesClient := factories.NewFactoriesClientWithBaseURI(endpoint)
	configureAuthFunc(&factoriesClient.Client)

	globalParametersClient := globalparameters.NewGlobalParametersClientWithBaseURI(endpoint)
	configureAuthFunc(&globalParametersClient.Client)

	integrationRuntimeNodesClient := integrationruntimenodes.NewIntegrationRuntimeNodesClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationRuntimeNodesClient.Client)

	integrationRuntimeObjectMetadataClient := integrationruntimeobjectmetadata.NewIntegrationRuntimeObjectMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationRuntimeObjectMetadataClient.Client)

	integrationRuntimesClient := integrationruntimes.NewIntegrationRuntimesClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationRuntimesClient.Client)

	linkedServicesClient := linkedservices.NewLinkedServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&linkedServicesClient.Client)

	managedPrivateEndpointsClient := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedPrivateEndpointsClient.Client)

	managedVirtualNetworksClient := managedvirtualnetworks.NewManagedVirtualNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&managedVirtualNetworksClient.Client)

	pipelineRunsClient := pipelineruns.NewPipelineRunsClientWithBaseURI(endpoint)
	configureAuthFunc(&pipelineRunsClient.Client)

	pipelinesClient := pipelines.NewPipelinesClientWithBaseURI(endpoint)
	configureAuthFunc(&pipelinesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	triggerClient := trigger.NewTriggerClientWithBaseURI(endpoint)
	configureAuthFunc(&triggerClient.Client)

	triggerrunsClient := triggerruns.NewTriggerrunsClientWithBaseURI(endpoint)
	configureAuthFunc(&triggerrunsClient.Client)

	triggersClient := triggers.NewTriggersClientWithBaseURI(endpoint)
	configureAuthFunc(&triggersClient.Client)

	return Client{
		Activityruns:                     &activityrunsClient,
		Credentials:                      &credentialsClient,
		DataFlowDebugSession:             &dataFlowDebugSessionClient,
		DataFlows:                        &dataFlowsClient,
		Datasets:                         &datasetsClient,
		ExposureControl:                  &exposureControlClient,
		Factories:                        &factoriesClient,
		GlobalParameters:                 &globalParametersClient,
		IntegrationRuntimeNodes:          &integrationRuntimeNodesClient,
		IntegrationRuntimeObjectMetadata: &integrationRuntimeObjectMetadataClient,
		IntegrationRuntimes:              &integrationRuntimesClient,
		LinkedServices:                   &linkedServicesClient,
		ManagedPrivateEndpoints:          &managedPrivateEndpointsClient,
		ManagedVirtualNetworks:           &managedVirtualNetworksClient,
		PipelineRuns:                     &pipelineRunsClient,
		Pipelines:                        &pipelinesClient,
		PrivateEndpointConnections:       &privateEndpointConnectionsClient,
		PrivateLinkResources:             &privateLinkResourcesClient,
		Trigger:                          &triggerClient,
		Triggerruns:                      &triggerrunsClient,
		Triggers:                         &triggersClient,
	}
}
