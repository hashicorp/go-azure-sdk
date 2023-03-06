// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2019_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/activity"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/agentregistrationinformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/automationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/certificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/connection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/connectiontype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/credential"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/dsccompilationjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/dscconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/dscnode"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/dscnodeconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/hybridrunbookworkergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/job"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/jobschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/jobstream"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/linkedworkspace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/listkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/module"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/nodecountinformation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/nodereports"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/objectdatatypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/python2package"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/runbook"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/runbookdraft"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/softwareupdateconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/softwareupdateconfigurationmachinerun"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/softwareupdateconfigurationrun"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/sourcecontrol"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/sourcecontrolsyncjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/sourcecontrolsyncjobstreams"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/statistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/testjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/testjobstream"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/typefields"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/variable"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2019-06-01/watcher"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Activity                              *activity.ActivityClient
	AgentRegistrationInformation          *agentregistrationinformation.AgentRegistrationInformationClient
	AutomationAccount                     *automationaccount.AutomationAccountClient
	Certificate                           *certificate.CertificateClient
	Connection                            *connection.ConnectionClient
	ConnectionType                        *connectiontype.ConnectionTypeClient
	Credential                            *credential.CredentialClient
	DscCompilationJob                     *dsccompilationjob.DscCompilationJobClient
	DscConfiguration                      *dscconfiguration.DscConfigurationClient
	DscNode                               *dscnode.DscNodeClient
	DscNodeConfiguration                  *dscnodeconfiguration.DscNodeConfigurationClient
	HybridRunbookWorkerGroup              *hybridrunbookworkergroup.HybridRunbookWorkerGroupClient
	Job                                   *job.JobClient
	JobSchedule                           *jobschedule.JobScheduleClient
	JobStream                             *jobstream.JobStreamClient
	LinkedWorkspace                       *linkedworkspace.LinkedWorkspaceClient
	ListKeys                              *listkeys.ListKeysClient
	Module                                *module.ModuleClient
	NodeCountInformation                  *nodecountinformation.NodeCountInformationClient
	NodeReports                           *nodereports.NodeReportsClient
	ObjectDataTypes                       *objectdatatypes.ObjectDataTypesClient
	Python2Package                        *python2package.Python2PackageClient
	Runbook                               *runbook.RunbookClient
	RunbookDraft                          *runbookdraft.RunbookDraftClient
	Schedule                              *schedule.ScheduleClient
	SoftwareUpdateConfiguration           *softwareupdateconfiguration.SoftwareUpdateConfigurationClient
	SoftwareUpdateConfigurationMachineRun *softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient
	SoftwareUpdateConfigurationRun        *softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient
	SourceControl                         *sourcecontrol.SourceControlClient
	SourceControlSyncJob                  *sourcecontrolsyncjob.SourceControlSyncJobClient
	SourceControlSyncJobStreams           *sourcecontrolsyncjobstreams.SourceControlSyncJobStreamsClient
	Statistics                            *statistics.StatisticsClient
	TestJob                               *testjob.TestJobClient
	TestJobStream                         *testjobstream.TestJobStreamClient
	TypeFields                            *typefields.TypeFieldsClient
	Usages                                *usages.UsagesClient
	Variable                              *variable.VariableClient
	Watcher                               *watcher.WatcherClient
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

	dscConfigurationClient := dscconfiguration.NewDscConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscConfigurationClient.Client)

	dscNodeClient := dscnode.NewDscNodeClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeClient.Client)

	dscNodeConfigurationClient := dscnodeconfiguration.NewDscNodeConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeConfigurationClient.Client)

	hybridRunbookWorkerGroupClient := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridRunbookWorkerGroupClient.Client)

	jobClient := job.NewJobClientWithBaseURI(endpoint)
	configureAuthFunc(&jobClient.Client)

	jobScheduleClient := jobschedule.NewJobScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&jobScheduleClient.Client)

	jobStreamClient := jobstream.NewJobStreamClientWithBaseURI(endpoint)
	configureAuthFunc(&jobStreamClient.Client)

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

	python2PackageClient := python2package.NewPython2PackageClientWithBaseURI(endpoint)
	configureAuthFunc(&python2PackageClient.Client)

	runbookClient := runbook.NewRunbookClientWithBaseURI(endpoint)
	configureAuthFunc(&runbookClient.Client)

	runbookDraftClient := runbookdraft.NewRunbookDraftClientWithBaseURI(endpoint)
	configureAuthFunc(&runbookDraftClient.Client)

	scheduleClient := schedule.NewScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduleClient.Client)

	softwareUpdateConfigurationClient := softwareupdateconfiguration.NewSoftwareUpdateConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&softwareUpdateConfigurationClient.Client)

	softwareUpdateConfigurationMachineRunClient := softwareupdateconfigurationmachinerun.NewSoftwareUpdateConfigurationMachineRunClientWithBaseURI(endpoint)
	configureAuthFunc(&softwareUpdateConfigurationMachineRunClient.Client)

	softwareUpdateConfigurationRunClient := softwareupdateconfigurationrun.NewSoftwareUpdateConfigurationRunClientWithBaseURI(endpoint)
	configureAuthFunc(&softwareUpdateConfigurationRunClient.Client)

	sourceControlClient := sourcecontrol.NewSourceControlClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlClient.Client)

	sourceControlSyncJobClient := sourcecontrolsyncjob.NewSourceControlSyncJobClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlSyncJobClient.Client)

	sourceControlSyncJobStreamsClient := sourcecontrolsyncjobstreams.NewSourceControlSyncJobStreamsClientWithBaseURI(endpoint)
	configureAuthFunc(&sourceControlSyncJobStreamsClient.Client)

	statisticsClient := statistics.NewStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&statisticsClient.Client)

	testJobClient := testjob.NewTestJobClientWithBaseURI(endpoint)
	configureAuthFunc(&testJobClient.Client)

	testJobStreamClient := testjobstream.NewTestJobStreamClientWithBaseURI(endpoint)
	configureAuthFunc(&testJobStreamClient.Client)

	typeFieldsClient := typefields.NewTypeFieldsClientWithBaseURI(endpoint)
	configureAuthFunc(&typeFieldsClient.Client)

	usagesClient := usages.NewUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&usagesClient.Client)

	variableClient := variable.NewVariableClientWithBaseURI(endpoint)
	configureAuthFunc(&variableClient.Client)

	watcherClient := watcher.NewWatcherClientWithBaseURI(endpoint)
	configureAuthFunc(&watcherClient.Client)

	return Client{
		Activity:                              &activityClient,
		AgentRegistrationInformation:          &agentRegistrationInformationClient,
		AutomationAccount:                     &automationAccountClient,
		Certificate:                           &certificateClient,
		Connection:                            &connectionClient,
		ConnectionType:                        &connectionTypeClient,
		Credential:                            &credentialClient,
		DscCompilationJob:                     &dscCompilationJobClient,
		DscConfiguration:                      &dscConfigurationClient,
		DscNode:                               &dscNodeClient,
		DscNodeConfiguration:                  &dscNodeConfigurationClient,
		HybridRunbookWorkerGroup:              &hybridRunbookWorkerGroupClient,
		Job:                                   &jobClient,
		JobSchedule:                           &jobScheduleClient,
		JobStream:                             &jobStreamClient,
		LinkedWorkspace:                       &linkedWorkspaceClient,
		ListKeys:                              &listKeysClient,
		Module:                                &moduleClient,
		NodeCountInformation:                  &nodeCountInformationClient,
		NodeReports:                           &nodeReportsClient,
		ObjectDataTypes:                       &objectDataTypesClient,
		Python2Package:                        &python2PackageClient,
		Runbook:                               &runbookClient,
		RunbookDraft:                          &runbookDraftClient,
		Schedule:                              &scheduleClient,
		SoftwareUpdateConfiguration:           &softwareUpdateConfigurationClient,
		SoftwareUpdateConfigurationMachineRun: &softwareUpdateConfigurationMachineRunClient,
		SoftwareUpdateConfigurationRun:        &softwareUpdateConfigurationRunClient,
		SourceControl:                         &sourceControlClient,
		SourceControlSyncJob:                  &sourceControlSyncJobClient,
		SourceControlSyncJobStreams:           &sourceControlSyncJobStreamsClient,
		Statistics:                            &statisticsClient,
		TestJob:                               &testJobClient,
		TestJobStream:                         &testJobStreamClient,
		TypeFields:                            &typeFieldsClient,
		Usages:                                &usagesClient,
		Variable:                              &variableClient,
		Watcher:                               &watcherClient,
	}
}
