package v2022_08_08

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	activityClient := activity.NewActivityClientWithBaseURI(endpoint)
	configureAuthFunc(&activityClient.Client)

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

	dscConfigurationClient := dscconfiguration.NewDscConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscConfigurationClient.Client)

	dscNodeConfigurationClient := dscnodeconfiguration.NewDscNodeConfigurationClientWithBaseURI(endpoint)
	configureAuthFunc(&dscNodeConfigurationClient.Client)

	hybridRunbookWorkerClient := hybridrunbookworker.NewHybridRunbookWorkerClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridRunbookWorkerClient.Client)

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

	listAllHybridRunbookWorkerGroupInAutomationAccountClient := listallhybridrunbookworkergroupinautomationaccount.NewListAllHybridRunbookWorkerGroupInAutomationAccountClientWithBaseURI(endpoint)
	configureAuthFunc(&listAllHybridRunbookWorkerGroupInAutomationAccountClient.Client)

	listKeysClient := listkeys.NewListKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&listKeysClient.Client)

	moduleClient := module.NewModuleClientWithBaseURI(endpoint)
	configureAuthFunc(&moduleClient.Client)

	objectDataTypesClient := objectdatatypes.NewObjectDataTypesClientWithBaseURI(endpoint)
	configureAuthFunc(&objectDataTypesClient.Client)

	python2PackageClient := python2package.NewPython2PackageClientWithBaseURI(endpoint)
	configureAuthFunc(&python2PackageClient.Client)

	python3PackageClient := python3package.NewPython3PackageClientWithBaseURI(endpoint)
	configureAuthFunc(&python3PackageClient.Client)

	runbookClient := runbook.NewRunbookClientWithBaseURI(endpoint)
	configureAuthFunc(&runbookClient.Client)

	runbookDraftClient := runbookdraft.NewRunbookDraftClientWithBaseURI(endpoint)
	configureAuthFunc(&runbookDraftClient.Client)

	scheduleClient := schedule.NewScheduleClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduleClient.Client)

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

	return Client{
		Activity:                 &activityClient,
		AutomationAccount:        &automationAccountClient,
		Certificate:              &certificateClient,
		Connection:               &connectionClient,
		ConnectionType:           &connectionTypeClient,
		Credential:               &credentialClient,
		DscConfiguration:         &dscConfigurationClient,
		DscNodeConfiguration:     &dscNodeConfigurationClient,
		HybridRunbookWorker:      &hybridRunbookWorkerClient,
		HybridRunbookWorkerGroup: &hybridRunbookWorkerGroupClient,
		Job:                      &jobClient,
		JobSchedule:              &jobScheduleClient,
		JobStream:                &jobStreamClient,
		LinkedWorkspace:          &linkedWorkspaceClient,
		ListAllHybridRunbookWorkerGroupInAutomationAccount: &listAllHybridRunbookWorkerGroupInAutomationAccountClient,
		ListKeys:                              &listKeysClient,
		Module:                                &moduleClient,
		ObjectDataTypes:                       &objectDataTypesClient,
		Python2Package:                        &python2PackageClient,
		Python3Package:                        &python3PackageClient,
		Runbook:                               &runbookClient,
		RunbookDraft:                          &runbookDraftClient,
		Schedule:                              &scheduleClient,
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
	}
}
