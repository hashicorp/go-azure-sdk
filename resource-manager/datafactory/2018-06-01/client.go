package v2018_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/activityruns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/changedatacapture"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/credentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/dataflowdebugsession"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/dataflows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/datasets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/exposurecontrol"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/factories"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/globalparameters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimedisableinteractivequery"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datafactory/2018-06-01/integrationruntimeenableinteractivequery"
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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Activityruns                              *activityruns.ActivityrunsClient
	ChangeDataCapture                         *changedatacapture.ChangeDataCaptureClient
	Credentials                               *credentials.CredentialsClient
	DataFlowDebugSession                      *dataflowdebugsession.DataFlowDebugSessionClient
	DataFlows                                 *dataflows.DataFlowsClient
	Datasets                                  *datasets.DatasetsClient
	ExposureControl                           *exposurecontrol.ExposureControlClient
	Factories                                 *factories.FactoriesClient
	GlobalParameters                          *globalparameters.GlobalParametersClient
	IntegrationRuntimeDisableInteractiveQuery *integrationruntimedisableinteractivequery.IntegrationRuntimeDisableInteractiveQueryClient
	IntegrationRuntimeEnableInteractiveQuery  *integrationruntimeenableinteractivequery.IntegrationRuntimeEnableInteractiveQueryClient
	IntegrationRuntimeNodes                   *integrationruntimenodes.IntegrationRuntimeNodesClient
	IntegrationRuntimeObjectMetadata          *integrationruntimeobjectmetadata.IntegrationRuntimeObjectMetadataClient
	IntegrationRuntimes                       *integrationruntimes.IntegrationRuntimesClient
	LinkedServices                            *linkedservices.LinkedServicesClient
	ManagedPrivateEndpoints                   *managedprivateendpoints.ManagedPrivateEndpointsClient
	ManagedVirtualNetworks                    *managedvirtualnetworks.ManagedVirtualNetworksClient
	PipelineRuns                              *pipelineruns.PipelineRunsClient
	Pipelines                                 *pipelines.PipelinesClient
	PrivateEndpointConnections                *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                      *privatelinkresources.PrivateLinkResourcesClient
	Trigger                                   *trigger.TriggerClient
	Triggerruns                               *triggerruns.TriggerrunsClient
	Triggers                                  *triggers.TriggersClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	activityrunsClient, err := activityruns.NewActivityrunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Activityruns client: %+v", err)
	}
	configureFunc(activityrunsClient.Client)

	changeDataCaptureClient, err := changedatacapture.NewChangeDataCaptureClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ChangeDataCapture client: %+v", err)
	}
	configureFunc(changeDataCaptureClient.Client)

	credentialsClient, err := credentials.NewCredentialsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Credentials client: %+v", err)
	}
	configureFunc(credentialsClient.Client)

	dataFlowDebugSessionClient, err := dataflowdebugsession.NewDataFlowDebugSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataFlowDebugSession client: %+v", err)
	}
	configureFunc(dataFlowDebugSessionClient.Client)

	dataFlowsClient, err := dataflows.NewDataFlowsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataFlows client: %+v", err)
	}
	configureFunc(dataFlowsClient.Client)

	datasetsClient, err := datasets.NewDatasetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Datasets client: %+v", err)
	}
	configureFunc(datasetsClient.Client)

	exposureControlClient, err := exposurecontrol.NewExposureControlClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ExposureControl client: %+v", err)
	}
	configureFunc(exposureControlClient.Client)

	factoriesClient, err := factories.NewFactoriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Factories client: %+v", err)
	}
	configureFunc(factoriesClient.Client)

	globalParametersClient, err := globalparameters.NewGlobalParametersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GlobalParameters client: %+v", err)
	}
	configureFunc(globalParametersClient.Client)

	integrationRuntimeDisableInteractiveQueryClient, err := integrationruntimedisableinteractivequery.NewIntegrationRuntimeDisableInteractiveQueryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimeDisableInteractiveQuery client: %+v", err)
	}
	configureFunc(integrationRuntimeDisableInteractiveQueryClient.Client)

	integrationRuntimeEnableInteractiveQueryClient, err := integrationruntimeenableinteractivequery.NewIntegrationRuntimeEnableInteractiveQueryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimeEnableInteractiveQuery client: %+v", err)
	}
	configureFunc(integrationRuntimeEnableInteractiveQueryClient.Client)

	integrationRuntimeNodesClient, err := integrationruntimenodes.NewIntegrationRuntimeNodesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimeNodes client: %+v", err)
	}
	configureFunc(integrationRuntimeNodesClient.Client)

	integrationRuntimeObjectMetadataClient, err := integrationruntimeobjectmetadata.NewIntegrationRuntimeObjectMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimeObjectMetadata client: %+v", err)
	}
	configureFunc(integrationRuntimeObjectMetadataClient.Client)

	integrationRuntimesClient, err := integrationruntimes.NewIntegrationRuntimesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimes client: %+v", err)
	}
	configureFunc(integrationRuntimesClient.Client)

	linkedServicesClient, err := linkedservices.NewLinkedServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedServices client: %+v", err)
	}
	configureFunc(linkedServicesClient.Client)

	managedPrivateEndpointsClient, err := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedPrivateEndpoints client: %+v", err)
	}
	configureFunc(managedPrivateEndpointsClient.Client)

	managedVirtualNetworksClient, err := managedvirtualnetworks.NewManagedVirtualNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedVirtualNetworks client: %+v", err)
	}
	configureFunc(managedVirtualNetworksClient.Client)

	pipelineRunsClient, err := pipelineruns.NewPipelineRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PipelineRuns client: %+v", err)
	}
	configureFunc(pipelineRunsClient.Client)

	pipelinesClient, err := pipelines.NewPipelinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pipelines client: %+v", err)
	}
	configureFunc(pipelinesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	triggerClient, err := trigger.NewTriggerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Trigger client: %+v", err)
	}
	configureFunc(triggerClient.Client)

	triggerrunsClient, err := triggerruns.NewTriggerrunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Triggerruns client: %+v", err)
	}
	configureFunc(triggerrunsClient.Client)

	triggersClient, err := triggers.NewTriggersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Triggers client: %+v", err)
	}
	configureFunc(triggersClient.Client)

	return &Client{
		Activityruns:         activityrunsClient,
		ChangeDataCapture:    changeDataCaptureClient,
		Credentials:          credentialsClient,
		DataFlowDebugSession: dataFlowDebugSessionClient,
		DataFlows:            dataFlowsClient,
		Datasets:             datasetsClient,
		ExposureControl:      exposureControlClient,
		Factories:            factoriesClient,
		GlobalParameters:     globalParametersClient,
		IntegrationRuntimeDisableInteractiveQuery: integrationRuntimeDisableInteractiveQueryClient,
		IntegrationRuntimeEnableInteractiveQuery:  integrationRuntimeEnableInteractiveQueryClient,
		IntegrationRuntimeNodes:                   integrationRuntimeNodesClient,
		IntegrationRuntimeObjectMetadata:          integrationRuntimeObjectMetadataClient,
		IntegrationRuntimes:                       integrationRuntimesClient,
		LinkedServices:                            linkedServicesClient,
		ManagedPrivateEndpoints:                   managedPrivateEndpointsClient,
		ManagedVirtualNetworks:                    managedVirtualNetworksClient,
		PipelineRuns:                              pipelineRunsClient,
		Pipelines:                                 pipelinesClient,
		PrivateEndpointConnections:                privateEndpointConnectionsClient,
		PrivateLinkResources:                      privateLinkResourcesClient,
		Trigger:                                   triggerClient,
		Triggerruns:                               triggerrunsClient,
		Triggers:                                  triggersClient,
	}, nil
}
