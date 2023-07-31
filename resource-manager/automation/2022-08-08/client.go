package v2022_08_08

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/activity"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/automationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/certificate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/connection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/connectiontype"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/credential"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/dscconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/dscnodeconfiguration"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/hybridrunbookworker"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/hybridrunbookworkergroup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/job"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/jobschedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/jobstream"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/linkedworkspace"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/listallhybridrunbookworkergroupinautomationaccount"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/listkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/module"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/objectdatatypes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/operations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/python2package"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/python3package"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/runbook"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/runbookdraft"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/schedule"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/softwareupdateconfigurationmachinerun"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/softwareupdateconfigurationrun"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/sourcecontrol"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/sourcecontrolsyncjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/sourcecontrolsyncjobstreams"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/statistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/testjob"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/testjobstream"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/typefields"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/automation/2022-08-08/variable"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Activity                                           *activity.ActivityClient
	AutomationAccount                                  *automationaccount.AutomationAccountClient
	Certificate                                        *certificate.CertificateClient
	Connection                                         *connection.ConnectionClient
	ConnectionType                                     *connectiontype.ConnectionTypeClient
	Credential                                         *credential.CredentialClient
	DscConfiguration                                   *dscconfiguration.DscConfigurationClient
	DscNodeConfiguration                               *dscnodeconfiguration.DscNodeConfigurationClient
	HybridRunbookWorker                                *hybridrunbookworker.HybridRunbookWorkerClient
	HybridRunbookWorkerGroup                           *hybridrunbookworkergroup.HybridRunbookWorkerGroupClient
	Job                                                *job.JobClient
	JobSchedule                                        *jobschedule.JobScheduleClient
	JobStream                                          *jobstream.JobStreamClient
	LinkedWorkspace                                    *linkedworkspace.LinkedWorkspaceClient
	ListAllHybridRunbookWorkerGroupInAutomationAccount *listallhybridrunbookworkergroupinautomationaccount.ListAllHybridRunbookWorkerGroupInAutomationAccountClient
	ListKeys                                           *listkeys.ListKeysClient
	Module                                             *module.ModuleClient
	ObjectDataTypes                                    *objectdatatypes.ObjectDataTypesClient
	Operations                                         *operations.OperationsClient
	Python2Package                                     *python2package.Python2PackageClient
	Python3Package                                     *python3package.Python3PackageClient
	Runbook                                            *runbook.RunbookClient
	RunbookDraft                                       *runbookdraft.RunbookDraftClient
	Schedule                                           *schedule.ScheduleClient
	SoftwareUpdateConfigurationMachineRun              *softwareupdateconfigurationmachinerun.SoftwareUpdateConfigurationMachineRunClient
	SoftwareUpdateConfigurationRun                     *softwareupdateconfigurationrun.SoftwareUpdateConfigurationRunClient
	SourceControl                                      *sourcecontrol.SourceControlClient
	SourceControlSyncJob                               *sourcecontrolsyncjob.SourceControlSyncJobClient
	SourceControlSyncJobStreams                        *sourcecontrolsyncjobstreams.SourceControlSyncJobStreamsClient
	Statistics                                         *statistics.StatisticsClient
	TestJob                                            *testjob.TestJobClient
	TestJobStream                                      *testjobstream.TestJobStreamClient
	TypeFields                                         *typefields.TypeFieldsClient
	Usages                                             *usages.UsagesClient
	Variable                                           *variable.VariableClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	activityClient, err := activity.NewActivityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Activity client: %+v", err)
	}
	configureFunc(activityClient.Client)

	automationAccountClient, err := automationaccount.NewAutomationAccountClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AutomationAccount client: %+v", err)
	}
	configureFunc(automationAccountClient.Client)

	certificateClient, err := certificate.NewCertificateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Certificate client: %+v", err)
	}
	configureFunc(certificateClient.Client)

	connectionClient, err := connection.NewConnectionClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Connection client: %+v", err)
	}
	configureFunc(connectionClient.Client)

	connectionTypeClient, err := connectiontype.NewConnectionTypeClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConnectionType client: %+v", err)
	}
	configureFunc(connectionTypeClient.Client)

	credentialClient, err := credential.NewCredentialClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Credential client: %+v", err)
	}
	configureFunc(credentialClient.Client)

	dscConfigurationClient, err := dscconfiguration.NewDscConfigurationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DscConfiguration client: %+v", err)
	}
	configureFunc(dscConfigurationClient.Client)

	dscNodeConfigurationClient, err := dscnodeconfiguration.NewDscNodeConfigurationClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DscNodeConfiguration client: %+v", err)
	}
	configureFunc(dscNodeConfigurationClient.Client)

	hybridRunbookWorkerClient, err := hybridrunbookworker.NewHybridRunbookWorkerClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HybridRunbookWorker client: %+v", err)
	}
	configureFunc(hybridRunbookWorkerClient.Client)

	hybridRunbookWorkerGroupClient, err := hybridrunbookworkergroup.NewHybridRunbookWorkerGroupClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building HybridRunbookWorkerGroup client: %+v", err)
	}
	configureFunc(hybridRunbookWorkerGroupClient.Client)

	jobClient, err := job.NewJobClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Job client: %+v", err)
	}
	configureFunc(jobClient.Client)

	jobScheduleClient, err := jobschedule.NewJobScheduleClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobSchedule client: %+v", err)
	}
	configureFunc(jobScheduleClient.Client)

	jobStreamClient, err := jobstream.NewJobStreamClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobStream client: %+v", err)
	}
	configureFunc(jobStreamClient.Client)

	linkedWorkspaceClient, err := linkedworkspace.NewLinkedWorkspaceClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LinkedWorkspace client: %+v", err)
	}
	configureFunc(linkedWorkspaceClient.Client)

	listAllHybridRunbookWorkerGroupInAutomationAccountClient, err := listallhybridrunbookworkergroupinautomationaccount.NewListAllHybridRunbookWorkerGroupInAutomationAccountClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ListAllHybridRunbookWorkerGroupInAutomationAccount client: %+v", err)
	}
	configureFunc(listAllHybridRunbookWorkerGroupInAutomationAccountClient.Client)

	listKeysClient, err := listkeys.NewListKeysClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ListKeys client: %+v", err)
	}
	configureFunc(listKeysClient.Client)

	moduleClient, err := module.NewModuleClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Module client: %+v", err)
	}
	configureFunc(moduleClient.Client)

	objectDataTypesClient, err := objectdatatypes.NewObjectDataTypesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ObjectDataTypes client: %+v", err)
	}
	configureFunc(objectDataTypesClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	python2PackageClient, err := python2package.NewPython2PackageClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Python2Package client: %+v", err)
	}
	configureFunc(python2PackageClient.Client)

	python3PackageClient, err := python3package.NewPython3PackageClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Python3Package client: %+v", err)
	}
	configureFunc(python3PackageClient.Client)

	runbookClient, err := runbook.NewRunbookClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Runbook client: %+v", err)
	}
	configureFunc(runbookClient.Client)

	runbookDraftClient, err := runbookdraft.NewRunbookDraftClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RunbookDraft client: %+v", err)
	}
	configureFunc(runbookDraftClient.Client)

	scheduleClient, err := schedule.NewScheduleClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Schedule client: %+v", err)
	}
	configureFunc(scheduleClient.Client)

	softwareUpdateConfigurationMachineRunClient, err := softwareupdateconfigurationmachinerun.NewSoftwareUpdateConfigurationMachineRunClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SoftwareUpdateConfigurationMachineRun client: %+v", err)
	}
	configureFunc(softwareUpdateConfigurationMachineRunClient.Client)

	softwareUpdateConfigurationRunClient, err := softwareupdateconfigurationrun.NewSoftwareUpdateConfigurationRunClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SoftwareUpdateConfigurationRun client: %+v", err)
	}
	configureFunc(softwareUpdateConfigurationRunClient.Client)

	sourceControlClient, err := sourcecontrol.NewSourceControlClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SourceControl client: %+v", err)
	}
	configureFunc(sourceControlClient.Client)

	sourceControlSyncJobClient, err := sourcecontrolsyncjob.NewSourceControlSyncJobClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SourceControlSyncJob client: %+v", err)
	}
	configureFunc(sourceControlSyncJobClient.Client)

	sourceControlSyncJobStreamsClient, err := sourcecontrolsyncjobstreams.NewSourceControlSyncJobStreamsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SourceControlSyncJobStreams client: %+v", err)
	}
	configureFunc(sourceControlSyncJobStreamsClient.Client)

	statisticsClient, err := statistics.NewStatisticsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Statistics client: %+v", err)
	}
	configureFunc(statisticsClient.Client)

	testJobClient, err := testjob.NewTestJobClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TestJob client: %+v", err)
	}
	configureFunc(testJobClient.Client)

	testJobStreamClient, err := testjobstream.NewTestJobStreamClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TestJobStream client: %+v", err)
	}
	configureFunc(testJobStreamClient.Client)

	typeFieldsClient, err := typefields.NewTypeFieldsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TypeFields client: %+v", err)
	}
	configureFunc(typeFieldsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	variableClient, err := variable.NewVariableClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Variable client: %+v", err)
	}
	configureFunc(variableClient.Client)

	return &Client{
		Activity:                 activityClient,
		AutomationAccount:        automationAccountClient,
		Certificate:              certificateClient,
		Connection:               connectionClient,
		ConnectionType:           connectionTypeClient,
		Credential:               credentialClient,
		DscConfiguration:         dscConfigurationClient,
		DscNodeConfiguration:     dscNodeConfigurationClient,
		HybridRunbookWorker:      hybridRunbookWorkerClient,
		HybridRunbookWorkerGroup: hybridRunbookWorkerGroupClient,
		Job:                      jobClient,
		JobSchedule:              jobScheduleClient,
		JobStream:                jobStreamClient,
		LinkedWorkspace:          linkedWorkspaceClient,
		ListAllHybridRunbookWorkerGroupInAutomationAccount: listAllHybridRunbookWorkerGroupInAutomationAccountClient,
		ListKeys:                              listKeysClient,
		Module:                                moduleClient,
		ObjectDataTypes:                       objectDataTypesClient,
		Operations:                            operationsClient,
		Python2Package:                        python2PackageClient,
		Python3Package:                        python3PackageClient,
		Runbook:                               runbookClient,
		RunbookDraft:                          runbookDraftClient,
		Schedule:                              scheduleClient,
		SoftwareUpdateConfigurationMachineRun: softwareUpdateConfigurationMachineRunClient,
		SoftwareUpdateConfigurationRun:        softwareUpdateConfigurationRunClient,
		SourceControl:                         sourceControlClient,
		SourceControlSyncJob:                  sourceControlSyncJobClient,
		SourceControlSyncJobStreams:           sourceControlSyncJobStreamsClient,
		Statistics:                            statisticsClient,
		TestJob:                               testJobClient,
		TestJobStream:                         testJobStreamClient,
		TypeFields:                            typeFieldsClient,
		Usages:                                usagesClient,
		Variable:                              variableClient,
	}, nil
}
