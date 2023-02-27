package v2020_01_13_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/activity"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/agentregistrationinformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/automationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/certificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/connection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/connectiontype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/credential"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/dsccompilationjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/dscnode"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/dscnodeconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/hybridrunbookworkergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/jobschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/linkedworkspace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/listkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/module"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/nodecountinformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/nodereports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/objectdatatypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/python2package"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/sourcecontrol"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/sourcecontrolsyncjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/sourcecontrolsyncjobstreams"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/statistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/typefields"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/variable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2020-01-13-preview/watcher"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Activity                     *activity.ActivityClient
	AgentRegistrationInformation *agentregistrationinformation.AgentRegistrationInformationClient
	AutomationAccount            *automationaccount.AutomationAccountClient
	Certificate                  *certificate.CertificateClient
	Connection                   *connection.ConnectionClient
	ConnectionType               *connectiontype.ConnectionTypeClient
	Credential                   *credential.CredentialClient
	DscCompilationJob            *dsccompilationjob.DscCompilationJobClient
	DscNode                      *dscnode.DscNodeClient
	DscNodeConfiguration         *dscnodeconfiguration.DscNodeConfigurationClient
	HybridRunbookWorkerGroup     *hybridrunbookworkergroup.HybridRunbookWorkerGroupClient
	JobSchedule                  *jobschedule.JobScheduleClient
	LinkedWorkspace              *linkedworkspace.LinkedWorkspaceClient
	ListKeys                     *listkeys.ListKeysClient
	Module                       *module.ModuleClient
	NodeCountInformation         *nodecountinformation.NodeCountInformationClient
	NodeReports                  *nodereports.NodeReportsClient
	ObjectDataTypes              *objectdatatypes.ObjectDataTypesClient
	PrivateEndpointConnections   *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources         *privatelinkresources.PrivateLinkResourcesClient
	Python2Package               *python2package.Python2PackageClient
	Schedule                     *schedule.ScheduleClient
	SourceControl                *sourcecontrol.SourceControlClient
	SourceControlSyncJob         *sourcecontrolsyncjob.SourceControlSyncJobClient
	SourceControlSyncJobStreams  *sourcecontrolsyncjobstreams.SourceControlSyncJobStreamsClient
	Statistics                   *statistics.StatisticsClient
	TypeFields                   *typefields.TypeFieldsClient
	Usages                       *usages.UsagesClient
	Variable                     *variable.VariableClient
	Watcher                      *watcher.WatcherClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	activityClient := activity.NewActivityClientWithBaseURI(endpoint)
	configureAuthFunc(&activityClient.Client)

	agentRegistrationInformationClient := agentregistrationinformation.NewAgentRegistrationInformationClientWithBaseURI(endpoint)
	configureAuthFunc(&agentRegistrationInformationClient.Client)

	automationAccountClient := automationaccount.NewAutomationAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&automationAccountClient.Client)

	certificateClient := certificate.NewCertificateClientWithBaseURI(endpoint)
	configureAuthFunc(&certificateClient.Client)

	connectionClient := connection.NewConnectionClientWithBaseURI(endpoint)
	configureAuthFunc(&connectionClient.Client)

	connectionTypeClient := connectiontype.NewConnectionTypeClientWithBaseURI(endpoint)
	configureAuthFunc(&connectionTypeClient.Client)

	credentialClient := credential.NewCredentialClientWithBaseURI(endpoint)
	configureAuthFunc(&credentialClient.Client)

	dscCompilationJobClient := dsccompilationjob.NewDscCompilationJobClientWithBaseURI(endpoint)
	configureAuthFunc(&dscCompilationJobClient.Client)

	dscNodeClient := dscnode.NewDscNodeClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeClient.Client)

	dscNodeConfigurationClient := dscnodeconfiguration.NewDscNodeConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeConfigurationClient.Client)

	hybridRunbookWorkerGroupClient := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridRunbookWorkerGroupClient.Client)

	jobScheduleClient := jobschedule.NewJobScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&jobScheduleClient.Client)

	linkedWorkspaceClient := linkedworkspace.NewLinkedWorkspaceClientWithBaseURI(endpoint)
	configureAuthFunc(&linkedWorkspaceClient.Client)

	listKeysClient := listkeys.NewListKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&listKeysClient.Client)

	moduleClient := module.NewModuleClientWithBaseURI(endpoint)
	configureAuthFunc(&moduleClient.Client)

	nodeCountInformationClient := nodecountinformation.NewNodeCountInformationClientWithBaseURI(endpoint)
	configureAuthFunc(&nodeCountInformationClient.Client)

	nodeReportsClient := nodereports.NewNodeReportsClientWithBaseURI(endpoint)
	configureAuthFunc(&nodeReportsClient.Client)

	objectDataTypesClient := objectdatatypes.NewObjectDataTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&objectDataTypesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	python2PackageClient := python2package.NewPython2PackageClientWithBaseURI(endpoint)
	configureAuthFunc(&python2PackageClient.Client)

	scheduleClient := schedule.NewScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduleClient.Client)

	sourceControlClient := sourcecontrol.NewSourceControlClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlClient.Client)

	sourceControlSyncJobClient := sourcecontrolsyncjob.NewSourceControlSyncJobClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlSyncJobClient.Client)

	sourceControlSyncJobStreamsClient := sourcecontrolsyncjobstreams.NewSourceControlSyncJobStreamsClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlSyncJobStreamsClient.Client)

	statisticsClient := statistics.NewStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&statisticsClient.Client)

	typeFieldsClient := typefields.NewTypeFieldsClientWithBaseURI(endpoint)
	configureAuthFunc(&typeFieldsClient.Client)

	usagesClient := usages.NewUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&usagesClient.Client)

	variableClient := variable.NewVariableClientWithBaseURI(endpoint)
	configureAuthFunc(&variableClient.Client)

	watcherClient := watcher.NewWatcherClientWithBaseURI(endpoint)
	configureAuthFunc(&watcherClient.Client)

	return Client{
		Activity:                     &activityClient,
		AgentRegistrationInformation: &agentRegistrationInformationClient,
		AutomationAccount:            &automationAccountClient,
		Certificate:                  &certificateClient,
		Connection:                   &connectionClient,
		ConnectionType:               &connectionTypeClient,
		Credential:                   &credentialClient,
		DscCompilationJob:            &dscCompilationJobClient,
		DscNode:                      &dscNodeClient,
		DscNodeConfiguration:         &dscNodeConfigurationClient,
		HybridRunbookWorkerGroup:     &hybridRunbookWorkerGroupClient,
		JobSchedule:                  &jobScheduleClient,
		LinkedWorkspace:              &linkedWorkspaceClient,
		ListKeys:                     &listKeysClient,
		Module:                       &moduleClient,
		NodeCountInformation:         &nodeCountInformationClient,
		NodeReports:                  &nodeReportsClient,
		ObjectDataTypes:              &objectDataTypesClient,
		PrivateEndpointConnections:   &privateEndpointConnectionsClient,
		PrivateLinkResources:         &privateLinkResourcesClient,
		Python2Package:               &python2PackageClient,
		Schedule:                     &scheduleClient,
		SourceControl:                &sourceControlClient,
		SourceControlSyncJob:         &sourceControlSyncJobClient,
		SourceControlSyncJobStreams:  &sourceControlSyncJobStreamsClient,
		Statistics:                   &statisticsClient,
		TypeFields:                   &typeFieldsClient,
		Usages:                       &usagesClient,
		Variable:                     &variableClient,
		Watcher:                      &watcherClient,
	}
}
