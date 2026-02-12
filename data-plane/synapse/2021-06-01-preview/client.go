package v2021_06_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/activityruns"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/bigdatapools"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/dataflowdebugsession"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/dataflows"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/datasets"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/integrationruntimes"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/kqlscripts"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/library"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/linkedservices"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/managedprivateendpoints"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/notebooks"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/operationresult"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/pipelineruns"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/pipelines"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/sparkconfigurations"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/sparkjobdefinitions"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/sqlpools"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/sqlscripts"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/trigger"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/triggerruns"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/triggers"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/workspace"
	"github.com/hashicorp/go-azure-sdk/data-plane/synapse/2021-06-01-preview/workspacegitrepomanagement"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

type Client struct {
	Activityruns               *activityruns.ActivityrunsClient
	BigDataPools               *bigdatapools.BigDataPoolsClient
	DataFlowDebugSession       *dataflowdebugsession.DataFlowDebugSessionClient
	DataFlows                  *dataflows.DataFlowsClient
	Datasets                   *datasets.DatasetsClient
	IntegrationRuntimes        *integrationruntimes.IntegrationRuntimesClient
	KqlScripts                 *kqlscripts.KqlScriptsClient
	Library                    *library.LibraryClient
	LinkedServices             *linkedservices.LinkedServicesClient
	ManagedPrivateEndpoints    *managedprivateendpoints.ManagedPrivateEndpointsClient
	Notebooks                  *notebooks.NotebooksClient
	OperationResult            *operationresult.OperationResultClient
	PipelineRuns               *pipelineruns.PipelineRunsClient
	Pipelines                  *pipelines.PipelinesClient
	SparkConfigurations        *sparkconfigurations.SparkConfigurationsClient
	SparkJobDefinitions        *sparkjobdefinitions.SparkJobDefinitionsClient
	SqlPools                   *sqlpools.SqlPoolsClient
	SqlScripts                 *sqlscripts.SqlScriptsClient
	Trigger                    *trigger.TriggerClient
	Triggerruns                *triggerruns.TriggerrunsClient
	Triggers                   *triggers.TriggersClient
	Workspace                  *workspace.WorkspaceClient
	WorkspaceGitRepoManagement *workspacegitrepomanagement.WorkspaceGitRepoManagementClient
}

func NewClient(configureFunc func(c *dataplane.Client)) (*Client, error) {
	activityrunsClient, err := activityruns.NewActivityrunsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Activityruns client: %+v", err)
	}
	configureFunc(activityrunsClient.Client)

	bigDataPoolsClient, err := bigdatapools.NewBigDataPoolsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building BigDataPools client: %+v", err)
	}
	configureFunc(bigDataPoolsClient.Client)

	dataFlowDebugSessionClient, err := dataflowdebugsession.NewDataFlowDebugSessionClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building DataFlowDebugSession client: %+v", err)
	}
	configureFunc(dataFlowDebugSessionClient.Client)

	dataFlowsClient, err := dataflows.NewDataFlowsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building DataFlows client: %+v", err)
	}
	configureFunc(dataFlowsClient.Client)

	datasetsClient, err := datasets.NewDatasetsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Datasets client: %+v", err)
	}
	configureFunc(datasetsClient.Client)

	integrationRuntimesClient, err := integrationruntimes.NewIntegrationRuntimesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimes client: %+v", err)
	}
	configureFunc(integrationRuntimesClient.Client)

	kqlScriptsClient, err := kqlscripts.NewKqlScriptsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building KqlScripts client: %+v", err)
	}
	configureFunc(kqlScriptsClient.Client)

	libraryClient, err := library.NewLibraryClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Library client: %+v", err)
	}
	configureFunc(libraryClient.Client)

	linkedServicesClient, err := linkedservices.NewLinkedServicesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building LinkedServices client: %+v", err)
	}
	configureFunc(linkedServicesClient.Client)

	managedPrivateEndpointsClient, err := managedprivateendpoints.NewManagedPrivateEndpointsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building ManagedPrivateEndpoints client: %+v", err)
	}
	configureFunc(managedPrivateEndpointsClient.Client)

	notebooksClient, err := notebooks.NewNotebooksClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Notebooks client: %+v", err)
	}
	configureFunc(notebooksClient.Client)

	operationResultClient, err := operationresult.NewOperationResultClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building OperationResult client: %+v", err)
	}
	configureFunc(operationResultClient.Client)

	pipelineRunsClient, err := pipelineruns.NewPipelineRunsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building PipelineRuns client: %+v", err)
	}
	configureFunc(pipelineRunsClient.Client)

	pipelinesClient, err := pipelines.NewPipelinesClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Pipelines client: %+v", err)
	}
	configureFunc(pipelinesClient.Client)

	sparkConfigurationsClient, err := sparkconfigurations.NewSparkConfigurationsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building SparkConfigurations client: %+v", err)
	}
	configureFunc(sparkConfigurationsClient.Client)

	sparkJobDefinitionsClient, err := sparkjobdefinitions.NewSparkJobDefinitionsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building SparkJobDefinitions client: %+v", err)
	}
	configureFunc(sparkJobDefinitionsClient.Client)

	sqlPoolsClient, err := sqlpools.NewSqlPoolsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building SqlPools client: %+v", err)
	}
	configureFunc(sqlPoolsClient.Client)

	sqlScriptsClient, err := sqlscripts.NewSqlScriptsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building SqlScripts client: %+v", err)
	}
	configureFunc(sqlScriptsClient.Client)

	triggerClient, err := trigger.NewTriggerClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Trigger client: %+v", err)
	}
	configureFunc(triggerClient.Client)

	triggerrunsClient, err := triggerruns.NewTriggerrunsClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Triggerruns client: %+v", err)
	}
	configureFunc(triggerrunsClient.Client)

	triggersClient, err := triggers.NewTriggersClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Triggers client: %+v", err)
	}
	configureFunc(triggersClient.Client)

	workspaceClient, err := workspace.NewWorkspaceClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building Workspace client: %+v", err)
	}
	configureFunc(workspaceClient.Client)

	workspaceGitRepoManagementClient, err := workspacegitrepomanagement.NewWorkspaceGitRepoManagementClientUnconfigured()
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceGitRepoManagement client: %+v", err)
	}
	configureFunc(workspaceGitRepoManagementClient.Client)

	return &Client{
		Activityruns:               activityrunsClient,
		BigDataPools:               bigDataPoolsClient,
		DataFlowDebugSession:       dataFlowDebugSessionClient,
		DataFlows:                  dataFlowsClient,
		Datasets:                   datasetsClient,
		IntegrationRuntimes:        integrationRuntimesClient,
		KqlScripts:                 kqlScriptsClient,
		Library:                    libraryClient,
		LinkedServices:             linkedServicesClient,
		ManagedPrivateEndpoints:    managedPrivateEndpointsClient,
		Notebooks:                  notebooksClient,
		OperationResult:            operationResultClient,
		PipelineRuns:               pipelineRunsClient,
		Pipelines:                  pipelinesClient,
		SparkConfigurations:        sparkConfigurationsClient,
		SparkJobDefinitions:        sparkJobDefinitionsClient,
		SqlPools:                   sqlPoolsClient,
		SqlScripts:                 sqlScriptsClient,
		Trigger:                    triggerClient,
		Triggerruns:                triggerrunsClient,
		Triggers:                   triggersClient,
		Workspace:                  workspaceClient,
		WorkspaceGitRepoManagement: workspaceGitRepoManagementClient,
	}, nil
}
