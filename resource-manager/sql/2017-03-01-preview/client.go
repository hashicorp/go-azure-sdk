package v2017_03_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/backuplongtermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/blobauditing"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/databasevulnerabilityassesmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/databasevulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/datawarehouseuseractivities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobagents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobcredentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobstepexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobsteps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobtargetexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobtargetgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/jobversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/longtermretentionbackups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/managedbackupshorttermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/manageddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/manageddatabasesecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/managedinstanceadministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/managedrestorabledroppeddatabasebackupshorttermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/managedserversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/restorabledroppedmanageddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/restorepoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/sensitivitylabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/serverautomatictuning"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/serverdnsaliases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2017-03-01-preview/serversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BackupLongTermRetentionPolicies                                  *backuplongtermretentionpolicies.BackupLongTermRetentionPoliciesClient
	BlobAuditing                                                     *blobauditing.BlobAuditingClient
	DataWarehouseUserActivities                                      *datawarehouseuseractivities.DataWarehouseUserActivitiesClient
	DatabaseVulnerabilityAssesmentRuleBaselines                      *databasevulnerabilityassesmentrulebaselines.DatabaseVulnerabilityAssesmentRuleBaselinesClient
	DatabaseVulnerabilityAssessments                                 *databasevulnerabilityassessments.DatabaseVulnerabilityAssessmentsClient
	Databases                                                        *databases.DatabasesClient
	JobAgents                                                        *jobagents.JobAgentsClient
	JobCredentials                                                   *jobcredentials.JobCredentialsClient
	JobExecutions                                                    *jobexecutions.JobExecutionsClient
	JobStepExecutions                                                *jobstepexecutions.JobStepExecutionsClient
	JobSteps                                                         *jobsteps.JobStepsClient
	JobTargetExecutions                                              *jobtargetexecutions.JobTargetExecutionsClient
	JobTargetGroups                                                  *jobtargetgroups.JobTargetGroupsClient
	JobVersions                                                      *jobversions.JobVersionsClient
	Jobs                                                             *jobs.JobsClient
	LongTermRetentionBackups                                         *longtermretentionbackups.LongTermRetentionBackupsClient
	ManagedBackupShortTermRetentionPolicies                          *managedbackupshorttermretentionpolicies.ManagedBackupShortTermRetentionPoliciesClient
	ManagedDatabaseSecurityAlertPolicies                             *manageddatabasesecurityalertpolicies.ManagedDatabaseSecurityAlertPoliciesClient
	ManagedDatabases                                                 *manageddatabases.ManagedDatabasesClient
	ManagedInstanceAdministrators                                    *managedinstanceadministrators.ManagedInstanceAdministratorsClient
	ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies *managedrestorabledroppeddatabasebackupshorttermretentionpolicies.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient
	ManagedServerSecurityAlertPolicies                               *managedserversecurityalertpolicies.ManagedServerSecurityAlertPoliciesClient
	RestorableDroppedManagedDatabases                                *restorabledroppedmanageddatabases.RestorableDroppedManagedDatabasesClient
	RestorePoints                                                    *restorepoints.RestorePointsClient
	SensitivityLabels                                                *sensitivitylabels.SensitivityLabelsClient
	ServerAutomaticTuning                                            *serverautomatictuning.ServerAutomaticTuningClient
	ServerDnsAliases                                                 *serverdnsaliases.ServerDnsAliasesClient
	ServerSecurityAlertPolicies                                      *serversecurityalertpolicies.ServerSecurityAlertPoliciesClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupLongTermRetentionPoliciesClient, err := backuplongtermretentionpolicies.NewBackupLongTermRetentionPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BackupLongTermRetentionPolicies client: %+v", err)
	}
	configureFunc(backupLongTermRetentionPoliciesClient.Client)

	blobAuditingClient, err := blobauditing.NewBlobAuditingClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building BlobAuditing client: %+v", err)
	}
	configureFunc(blobAuditingClient.Client)

	dataWarehouseUserActivitiesClient, err := datawarehouseuseractivities.NewDataWarehouseUserActivitiesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DataWarehouseUserActivities client: %+v", err)
	}
	configureFunc(dataWarehouseUserActivitiesClient.Client)

	databaseVulnerabilityAssesmentRuleBaselinesClient, err := databasevulnerabilityassesmentrulebaselines.NewDatabaseVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseVulnerabilityAssesmentRuleBaselines client: %+v", err)
	}
	configureFunc(databaseVulnerabilityAssesmentRuleBaselinesClient.Client)

	databaseVulnerabilityAssessmentsClient, err := databasevulnerabilityassessments.NewDatabaseVulnerabilityAssessmentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(databaseVulnerabilityAssessmentsClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	jobAgentsClient, err := jobagents.NewJobAgentsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobAgents client: %+v", err)
	}
	configureFunc(jobAgentsClient.Client)

	jobCredentialsClient, err := jobcredentials.NewJobCredentialsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobCredentials client: %+v", err)
	}
	configureFunc(jobCredentialsClient.Client)

	jobExecutionsClient, err := jobexecutions.NewJobExecutionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobExecutions client: %+v", err)
	}
	configureFunc(jobExecutionsClient.Client)

	jobStepExecutionsClient, err := jobstepexecutions.NewJobStepExecutionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobStepExecutions client: %+v", err)
	}
	configureFunc(jobStepExecutionsClient.Client)

	jobStepsClient, err := jobsteps.NewJobStepsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobSteps client: %+v", err)
	}
	configureFunc(jobStepsClient.Client)

	jobTargetExecutionsClient, err := jobtargetexecutions.NewJobTargetExecutionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobTargetExecutions client: %+v", err)
	}
	configureFunc(jobTargetExecutionsClient.Client)

	jobTargetGroupsClient, err := jobtargetgroups.NewJobTargetGroupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobTargetGroups client: %+v", err)
	}
	configureFunc(jobTargetGroupsClient.Client)

	jobVersionsClient, err := jobversions.NewJobVersionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building JobVersions client: %+v", err)
	}
	configureFunc(jobVersionsClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	longTermRetentionBackupsClient, err := longtermretentionbackups.NewLongTermRetentionBackupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LongTermRetentionBackups client: %+v", err)
	}
	configureFunc(longTermRetentionBackupsClient.Client)

	managedBackupShortTermRetentionPoliciesClient, err := managedbackupshorttermretentionpolicies.NewManagedBackupShortTermRetentionPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedBackupShortTermRetentionPolicies client: %+v", err)
	}
	configureFunc(managedBackupShortTermRetentionPoliciesClient.Client)

	managedDatabaseSecurityAlertPoliciesClient, err := manageddatabasesecurityalertpolicies.NewManagedDatabaseSecurityAlertPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(managedDatabaseSecurityAlertPoliciesClient.Client)

	managedDatabasesClient, err := manageddatabases.NewManagedDatabasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabases client: %+v", err)
	}
	configureFunc(managedDatabasesClient.Client)

	managedInstanceAdministratorsClient, err := managedinstanceadministrators.NewManagedInstanceAdministratorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceAdministrators client: %+v", err)
	}
	configureFunc(managedInstanceAdministratorsClient.Client)

	managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient, err := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies client: %+v", err)
	}
	configureFunc(managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Client)

	managedServerSecurityAlertPoliciesClient, err := managedserversecurityalertpolicies.NewManagedServerSecurityAlertPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ManagedServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(managedServerSecurityAlertPoliciesClient.Client)

	restorableDroppedManagedDatabasesClient, err := restorabledroppedmanageddatabases.NewRestorableDroppedManagedDatabasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RestorableDroppedManagedDatabases client: %+v", err)
	}
	configureFunc(restorableDroppedManagedDatabasesClient.Client)

	restorePointsClient, err := restorepoints.NewRestorePointsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RestorePoints client: %+v", err)
	}
	configureFunc(restorePointsClient.Client)

	sensitivityLabelsClient, err := sensitivitylabels.NewSensitivityLabelsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SensitivityLabels client: %+v", err)
	}
	configureFunc(sensitivityLabelsClient.Client)

	serverAutomaticTuningClient, err := serverautomatictuning.NewServerAutomaticTuningClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerAutomaticTuning client: %+v", err)
	}
	configureFunc(serverAutomaticTuningClient.Client)

	serverDnsAliasesClient, err := serverdnsaliases.NewServerDnsAliasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerDnsAliases client: %+v", err)
	}
	configureFunc(serverDnsAliasesClient.Client)

	serverSecurityAlertPoliciesClient, err := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(serverSecurityAlertPoliciesClient.Client)

	return &Client{
		BackupLongTermRetentionPolicies:             backupLongTermRetentionPoliciesClient,
		BlobAuditing:                                blobAuditingClient,
		DataWarehouseUserActivities:                 dataWarehouseUserActivitiesClient,
		DatabaseVulnerabilityAssesmentRuleBaselines: databaseVulnerabilityAssesmentRuleBaselinesClient,
		DatabaseVulnerabilityAssessments:            databaseVulnerabilityAssessmentsClient,
		Databases:                                   databasesClient,
		JobAgents:                                   jobAgentsClient,
		JobCredentials:                              jobCredentialsClient,
		JobExecutions:                               jobExecutionsClient,
		JobStepExecutions:                           jobStepExecutionsClient,
		JobSteps:                                    jobStepsClient,
		JobTargetExecutions:                         jobTargetExecutionsClient,
		JobTargetGroups:                             jobTargetGroupsClient,
		JobVersions:                                 jobVersionsClient,
		Jobs:                                        jobsClient,
		LongTermRetentionBackups:                    longTermRetentionBackupsClient,
		ManagedBackupShortTermRetentionPolicies:     managedBackupShortTermRetentionPoliciesClient,
		ManagedDatabaseSecurityAlertPolicies:        managedDatabaseSecurityAlertPoliciesClient,
		ManagedDatabases:                            managedDatabasesClient,
		ManagedInstanceAdministrators:               managedInstanceAdministratorsClient,
		ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies: managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient,
		ManagedServerSecurityAlertPolicies:                               managedServerSecurityAlertPoliciesClient,
		RestorableDroppedManagedDatabases:                                restorableDroppedManagedDatabasesClient,
		RestorePoints:                                                    restorePointsClient,
		SensitivityLabels:                                                sensitivityLabelsClient,
		ServerAutomaticTuning:                                            serverAutomaticTuningClient,
		ServerDnsAliases:                                                 serverDnsAliasesClient,
		ServerSecurityAlertPolicies:                                      serverSecurityAlertPoliciesClient,
	}, nil
}
