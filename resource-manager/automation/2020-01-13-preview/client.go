package v2020_01_13_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	activityClient, err := activity.NewActivityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Activity client: %+v", err)
	}
	configureFunc(activityClient.Client)

	agentRegistrationInformationClient, err := agentregistrationinformation.NewAgentRegistrationInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AgentRegistrationInformation client: %+v", err)
	}
	configureFunc(agentRegistrationInformationClient.Client)

	automationAccountClient, err := automationaccount.NewAutomationAccountClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutomationAccount client: %+v", err)
	}
	configureFunc(automationAccountClient.Client)

	certificateClient, err := certificate.NewCertificateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Certificate client: %+v", err)
	}
	configureFunc(certificateClient.Client)

	connectionClient, err := connection.NewConnectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Connection client: %+v", err)
	}
	configureFunc(connectionClient.Client)

	connectionTypeClient, err := connectiontype.NewConnectionTypeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectionType client: %+v", err)
	}
	configureFunc(connectionTypeClient.Client)

	credentialClient, err := credential.NewCredentialClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Credential client: %+v", err)
	}
	configureFunc(credentialClient.Client)

	dscCompilationJobClient, err := dsccompilationjob.NewDscCompilationJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DscCompilationJob client: %+v", err)
	}
	configureFunc(dscCompilationJobClient.Client)

	dscNodeClient, err := dscnode.NewDscNodeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DscNode client: %+v", err)
	}
	configureFunc(dscNodeClient.Client)

	dscNodeConfigurationClient, err := dscnodeconfiguration.NewDscNodeConfigurationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DscNodeConfiguration client: %+v", err)
	}
	configureFunc(dscNodeConfigurationClient.Client)

	hybridRunbookWorkerGroupClient, err := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridRunbookWorkerGroup client: %+v", err)
	}
	configureFunc(hybridRunbookWorkerGroupClient.Client)

	jobScheduleClient, err := jobschedule.NewJobScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobSchedule client: %+v", err)
	}
	configureFunc(jobScheduleClient.Client)

	linkedWorkspaceClient, err := linkedworkspace.NewLinkedWorkspaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LinkedWorkspace client: %+v", err)
	}
	configureFunc(linkedWorkspaceClient.Client)

	listKeysClient, err := listkeys.NewListKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ListKeys client: %+v", err)
	}
	configureFunc(listKeysClient.Client)

	moduleClient, err := module.NewModuleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Module client: %+v", err)
	}
	configureFunc(moduleClient.Client)

	nodeCountInformationClient, err := nodecountinformation.NewNodeCountInformationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NodeCountInformation client: %+v", err)
	}
	configureFunc(nodeCountInformationClient.Client)

	nodeReportsClient, err := nodereports.NewNodeReportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NodeReports client: %+v", err)
	}
	configureFunc(nodeReportsClient.Client)

	objectDataTypesClient, err := objectdatatypes.NewObjectDataTypesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ObjectDataTypes client: %+v", err)
	}
	configureFunc(objectDataTypesClient.Client)

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

	python2PackageClient, err := python2package.NewPython2PackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Python2Package client: %+v", err)
	}
	configureFunc(python2PackageClient.Client)

	scheduleClient, err := schedule.NewScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %+v", err)
	}
	configureFunc(scheduleClient.Client)

	sourceControlClient, err := sourcecontrol.NewSourceControlClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SourceControl client: %+v", err)
	}
	configureFunc(sourceControlClient.Client)

	sourceControlSyncJobClient, err := sourcecontrolsyncjob.NewSourceControlSyncJobClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SourceControlSyncJob client: %+v", err)
	}
	configureFunc(sourceControlSyncJobClient.Client)

	sourceControlSyncJobStreamsClient, err := sourcecontrolsyncjobstreams.NewSourceControlSyncJobStreamsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SourceControlSyncJobStreams client: %+v", err)
	}
	configureFunc(sourceControlSyncJobStreamsClient.Client)

	statisticsClient, err := statistics.NewStatisticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Statistics client: %+v", err)
	}
	configureFunc(statisticsClient.Client)

	typeFieldsClient, err := typefields.NewTypeFieldsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TypeFields client: %+v", err)
	}
	configureFunc(typeFieldsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	variableClient, err := variable.NewVariableClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Variable client: %+v", err)
	}
	configureFunc(variableClient.Client)

	watcherClient, err := watcher.NewWatcherClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Watcher client: %+v", err)
	}
	configureFunc(watcherClient.Client)

	return &Client{
		Activity:                     activityClient,
		AgentRegistrationInformation: agentRegistrationInformationClient,
		AutomationAccount:            automationAccountClient,
		Certificate:                  certificateClient,
		Connection:                   connectionClient,
		ConnectionType:               connectionTypeClient,
		Credential:                   credentialClient,
		DscCompilationJob:            dscCompilationJobClient,
		DscNode:                      dscNodeClient,
		DscNodeConfiguration:         dscNodeConfigurationClient,
		HybridRunbookWorkerGroup:     hybridRunbookWorkerGroupClient,
		JobSchedule:                  jobScheduleClient,
		LinkedWorkspace:              linkedWorkspaceClient,
		ListKeys:                     listKeysClient,
		Module:                       moduleClient,
		NodeCountInformation:         nodeCountInformationClient,
		NodeReports:                  nodeReportsClient,
		ObjectDataTypes:              objectDataTypesClient,
		PrivateEndpointConnections:   privateEndpointConnectionsClient,
		PrivateLinkResources:         privateLinkResourcesClient,
		Python2Package:               python2PackageClient,
		Schedule:                     scheduleClient,
		SourceControl:                sourceControlClient,
		SourceControlSyncJob:         sourceControlSyncJobClient,
		SourceControlSyncJobStreams:  sourceControlSyncJobStreamsClient,
		Statistics:                   statisticsClient,
		TypeFields:                   typeFieldsClient,
		Usages:                       usagesClient,
		Variable:                     variableClient,
		Watcher:                      watcherClient,
	}, nil
}
