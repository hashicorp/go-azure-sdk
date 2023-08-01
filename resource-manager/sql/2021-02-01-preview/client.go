package v2021_02_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/backupshorttermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/blobauditing"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaseadvisors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaseautomatictuning"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databasecolumns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaseextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaseoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaserecommendedactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaseschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databasesecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databasetables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databaseusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databasevulnerabilityassessmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databasevulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/databasevulnerabilityassessmentscans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/datawarehouseuseractivities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/deletedservers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/elasticpooloperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/elasticpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/encryptionprotectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/failovergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/instancefailovergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/instancepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobagents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobcredentials"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobstepexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobsteps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobtargetexecutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobtargetgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/jobversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/ledgerdigestuploads"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/locationcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/longtermretentionbackups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/longtermretentionmanagedinstancebackups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/longtermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/maintenancewindowoptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/maintenancewindows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedbackupshorttermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasecolumns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasequeries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabaserestoredetails"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabaseschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesecurityevents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasesensitivitylabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasetables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasetransparentdataencryption"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasevulnerabilityassessmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasevulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/manageddatabasevulnerabilityassessmentscans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstanceadministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstanceazureadonlyauthentications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstanceencryptionprotectors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstancekeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstancelongtermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstanceoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstanceprivateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstanceprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstancetdecertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedinstancevulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedrestorabledroppeddatabasebackupshorttermretentionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/managedserversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/outboundfirewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/recoverablemanageddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/replicationlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/restorabledroppeddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/restorabledroppedmanageddatabases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/restorepoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/sensitivitylabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serveradvisors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serverautomatictuning"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serverazureadadministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serverazureadonlyauthentications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serverdevopsaudit"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serverdnsaliases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serverkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serveroperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/serversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/servertrustgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/servervulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/sqlagent"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/subscriptionusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/syncagents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/syncgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/syncmembers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/tdecertificates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/timezones"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/transparentdataencryptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/usages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/virtualclusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/virtualnetworkrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/workloadclassifiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/sql/2021-02-01-preview/workloadgroups"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BackupShortTermRetentionPolicies                                 *backupshorttermretentionpolicies.BackupShortTermRetentionPoliciesClient
	BlobAuditing                                                     *blobauditing.BlobAuditingClient
	DataWarehouseUserActivities                                      *datawarehouseuseractivities.DataWarehouseUserActivitiesClient
	DatabaseAdvisors                                                 *databaseadvisors.DatabaseAdvisorsClient
	DatabaseAutomaticTuning                                          *databaseautomatictuning.DatabaseAutomaticTuningClient
	DatabaseColumns                                                  *databasecolumns.DatabaseColumnsClient
	DatabaseExtensions                                               *databaseextensions.DatabaseExtensionsClient
	DatabaseOperations                                               *databaseoperations.DatabaseOperationsClient
	DatabaseRecommendedActions                                       *databaserecommendedactions.DatabaseRecommendedActionsClient
	DatabaseSchemas                                                  *databaseschemas.DatabaseSchemasClient
	DatabaseSecurityAlertPolicies                                    *databasesecurityalertpolicies.DatabaseSecurityAlertPoliciesClient
	DatabaseTables                                                   *databasetables.DatabaseTablesClient
	DatabaseUsages                                                   *databaseusages.DatabaseUsagesClient
	DatabaseVulnerabilityAssessmentRuleBaselines                     *databasevulnerabilityassessmentrulebaselines.DatabaseVulnerabilityAssessmentRuleBaselinesClient
	DatabaseVulnerabilityAssessmentScans                             *databasevulnerabilityassessmentscans.DatabaseVulnerabilityAssessmentScansClient
	DatabaseVulnerabilityAssessments                                 *databasevulnerabilityassessments.DatabaseVulnerabilityAssessmentsClient
	Databases                                                        *databases.DatabasesClient
	DeletedServers                                                   *deletedservers.DeletedServersClient
	ElasticPoolOperations                                            *elasticpooloperations.ElasticPoolOperationsClient
	ElasticPools                                                     *elasticpools.ElasticPoolsClient
	EncryptionProtectors                                             *encryptionprotectors.EncryptionProtectorsClient
	FailoverGroups                                                   *failovergroups.FailoverGroupsClient
	FirewallRules                                                    *firewallrules.FirewallRulesClient
	InstanceFailoverGroups                                           *instancefailovergroups.InstanceFailoverGroupsClient
	InstancePools                                                    *instancepools.InstancePoolsClient
	JobAgents                                                        *jobagents.JobAgentsClient
	JobCredentials                                                   *jobcredentials.JobCredentialsClient
	JobExecutions                                                    *jobexecutions.JobExecutionsClient
	JobStepExecutions                                                *jobstepexecutions.JobStepExecutionsClient
	JobSteps                                                         *jobsteps.JobStepsClient
	JobTargetExecutions                                              *jobtargetexecutions.JobTargetExecutionsClient
	JobTargetGroups                                                  *jobtargetgroups.JobTargetGroupsClient
	JobVersions                                                      *jobversions.JobVersionsClient
	Jobs                                                             *jobs.JobsClient
	LedgerDigestUploads                                              *ledgerdigestuploads.LedgerDigestUploadsClient
	LocationCapabilities                                             *locationcapabilities.LocationCapabilitiesClient
	LongTermRetentionBackups                                         *longtermretentionbackups.LongTermRetentionBackupsClient
	LongTermRetentionManagedInstanceBackups                          *longtermretentionmanagedinstancebackups.LongTermRetentionManagedInstanceBackupsClient
	LongTermRetentionPolicies                                        *longtermretentionpolicies.LongTermRetentionPoliciesClient
	MaintenanceWindowOptions                                         *maintenancewindowoptions.MaintenanceWindowOptionsClient
	MaintenanceWindows                                               *maintenancewindows.MaintenanceWindowsClient
	ManagedBackupShortTermRetentionPolicies                          *managedbackupshorttermretentionpolicies.ManagedBackupShortTermRetentionPoliciesClient
	ManagedDatabaseColumns                                           *manageddatabasecolumns.ManagedDatabaseColumnsClient
	ManagedDatabaseQueries                                           *manageddatabasequeries.ManagedDatabaseQueriesClient
	ManagedDatabaseRestoreDetails                                    *manageddatabaserestoredetails.ManagedDatabaseRestoreDetailsClient
	ManagedDatabaseSchemas                                           *manageddatabaseschemas.ManagedDatabaseSchemasClient
	ManagedDatabaseSecurityAlertPolicies                             *manageddatabasesecurityalertpolicies.ManagedDatabaseSecurityAlertPoliciesClient
	ManagedDatabaseSecurityEvents                                    *manageddatabasesecurityevents.ManagedDatabaseSecurityEventsClient
	ManagedDatabaseSensitivityLabels                                 *manageddatabasesensitivitylabels.ManagedDatabaseSensitivityLabelsClient
	ManagedDatabaseTables                                            *manageddatabasetables.ManagedDatabaseTablesClient
	ManagedDatabaseTransparentDataEncryption                         *manageddatabasetransparentdataencryption.ManagedDatabaseTransparentDataEncryptionClient
	ManagedDatabaseVulnerabilityAssessmentRuleBaselines              *manageddatabasevulnerabilityassessmentrulebaselines.ManagedDatabaseVulnerabilityAssessmentRuleBaselinesClient
	ManagedDatabaseVulnerabilityAssessmentScans                      *manageddatabasevulnerabilityassessmentscans.ManagedDatabaseVulnerabilityAssessmentScansClient
	ManagedDatabaseVulnerabilityAssessments                          *manageddatabasevulnerabilityassessments.ManagedDatabaseVulnerabilityAssessmentsClient
	ManagedDatabases                                                 *manageddatabases.ManagedDatabasesClient
	ManagedInstanceAdministrators                                    *managedinstanceadministrators.ManagedInstanceAdministratorsClient
	ManagedInstanceAzureADOnlyAuthentications                        *managedinstanceazureadonlyauthentications.ManagedInstanceAzureADOnlyAuthenticationsClient
	ManagedInstanceEncryptionProtectors                              *managedinstanceencryptionprotectors.ManagedInstanceEncryptionProtectorsClient
	ManagedInstanceKeys                                              *managedinstancekeys.ManagedInstanceKeysClient
	ManagedInstanceLongTermRetentionPolicies                         *managedinstancelongtermretentionpolicies.ManagedInstanceLongTermRetentionPoliciesClient
	ManagedInstanceOperations                                        *managedinstanceoperations.ManagedInstanceOperationsClient
	ManagedInstancePrivateEndpointConnections                        *managedinstanceprivateendpointconnections.ManagedInstancePrivateEndpointConnectionsClient
	ManagedInstancePrivateLinkResources                              *managedinstanceprivatelinkresources.ManagedInstancePrivateLinkResourcesClient
	ManagedInstanceTdeCertificates                                   *managedinstancetdecertificates.ManagedInstanceTdeCertificatesClient
	ManagedInstanceVulnerabilityAssessments                          *managedinstancevulnerabilityassessments.ManagedInstanceVulnerabilityAssessmentsClient
	ManagedInstances                                                 *managedinstances.ManagedInstancesClient
	ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies *managedrestorabledroppeddatabasebackupshorttermretentionpolicies.ManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient
	ManagedServerSecurityAlertPolicies                               *managedserversecurityalertpolicies.ManagedServerSecurityAlertPoliciesClient
	OutboundFirewallRules                                            *outboundfirewallrules.OutboundFirewallRulesClient
	PrivateEndpointConnections                                       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                                             *privatelinkresources.PrivateLinkResourcesClient
	RecoverableManagedDatabases                                      *recoverablemanageddatabases.RecoverableManagedDatabasesClient
	ReplicationLinks                                                 *replicationlinks.ReplicationLinksClient
	RestorableDroppedDatabases                                       *restorabledroppeddatabases.RestorableDroppedDatabasesClient
	RestorableDroppedManagedDatabases                                *restorabledroppedmanageddatabases.RestorableDroppedManagedDatabasesClient
	RestorePoints                                                    *restorepoints.RestorePointsClient
	SensitivityLabels                                                *sensitivitylabels.SensitivityLabelsClient
	ServerAdvisors                                                   *serveradvisors.ServerAdvisorsClient
	ServerAutomaticTuning                                            *serverautomatictuning.ServerAutomaticTuningClient
	ServerAzureADAdministrators                                      *serverazureadadministrators.ServerAzureADAdministratorsClient
	ServerAzureADOnlyAuthentications                                 *serverazureadonlyauthentications.ServerAzureADOnlyAuthenticationsClient
	ServerDevOpsAudit                                                *serverdevopsaudit.ServerDevOpsAuditClient
	ServerDnsAliases                                                 *serverdnsaliases.ServerDnsAliasesClient
	ServerKeys                                                       *serverkeys.ServerKeysClient
	ServerOperations                                                 *serveroperations.ServerOperationsClient
	ServerSecurityAlertPolicies                                      *serversecurityalertpolicies.ServerSecurityAlertPoliciesClient
	ServerTrustGroups                                                *servertrustgroups.ServerTrustGroupsClient
	ServerVulnerabilityAssessments                                   *servervulnerabilityassessments.ServerVulnerabilityAssessmentsClient
	Servers                                                          *servers.ServersClient
	SqlAgent                                                         *sqlagent.SqlAgentClient
	SubscriptionUsages                                               *subscriptionusages.SubscriptionUsagesClient
	SyncAgents                                                       *syncagents.SyncAgentsClient
	SyncGroups                                                       *syncgroups.SyncGroupsClient
	SyncMembers                                                      *syncmembers.SyncMembersClient
	TdeCertificates                                                  *tdecertificates.TdeCertificatesClient
	TimeZones                                                        *timezones.TimeZonesClient
	TransparentDataEncryptions                                       *transparentdataencryptions.TransparentDataEncryptionsClient
	Usages                                                           *usages.UsagesClient
	VirtualClusters                                                  *virtualclusters.VirtualClustersClient
	VirtualNetworkRules                                              *virtualnetworkrules.VirtualNetworkRulesClient
	WorkloadClassifiers                                              *workloadclassifiers.WorkloadClassifiersClient
	WorkloadGroups                                                   *workloadgroups.WorkloadGroupsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupShortTermRetentionPoliciesClient, err := backupshorttermretentionpolicies.NewBackupShortTermRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupShortTermRetentionPolicies client: %+v", err)
	}
	configureFunc(backupShortTermRetentionPoliciesClient.Client)

	blobAuditingClient, err := blobauditing.NewBlobAuditingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BlobAuditing client: %+v", err)
	}
	configureFunc(blobAuditingClient.Client)

	dataWarehouseUserActivitiesClient, err := datawarehouseuseractivities.NewDataWarehouseUserActivitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataWarehouseUserActivities client: %+v", err)
	}
	configureFunc(dataWarehouseUserActivitiesClient.Client)

	databaseAdvisorsClient, err := databaseadvisors.NewDatabaseAdvisorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseAdvisors client: %+v", err)
	}
	configureFunc(databaseAdvisorsClient.Client)

	databaseAutomaticTuningClient, err := databaseautomatictuning.NewDatabaseAutomaticTuningClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseAutomaticTuning client: %+v", err)
	}
	configureFunc(databaseAutomaticTuningClient.Client)

	databaseColumnsClient, err := databasecolumns.NewDatabaseColumnsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseColumns client: %+v", err)
	}
	configureFunc(databaseColumnsClient.Client)

	databaseExtensionsClient, err := databaseextensions.NewDatabaseExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseExtensions client: %+v", err)
	}
	configureFunc(databaseExtensionsClient.Client)

	databaseOperationsClient, err := databaseoperations.NewDatabaseOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseOperations client: %+v", err)
	}
	configureFunc(databaseOperationsClient.Client)

	databaseRecommendedActionsClient, err := databaserecommendedactions.NewDatabaseRecommendedActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseRecommendedActions client: %+v", err)
	}
	configureFunc(databaseRecommendedActionsClient.Client)

	databaseSchemasClient, err := databaseschemas.NewDatabaseSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseSchemas client: %+v", err)
	}
	configureFunc(databaseSchemasClient.Client)

	databaseSecurityAlertPoliciesClient, err := databasesecurityalertpolicies.NewDatabaseSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(databaseSecurityAlertPoliciesClient.Client)

	databaseTablesClient, err := databasetables.NewDatabaseTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseTables client: %+v", err)
	}
	configureFunc(databaseTablesClient.Client)

	databaseUsagesClient, err := databaseusages.NewDatabaseUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseUsages client: %+v", err)
	}
	configureFunc(databaseUsagesClient.Client)

	databaseVulnerabilityAssessmentRuleBaselinesClient, err := databasevulnerabilityassessmentrulebaselines.NewDatabaseVulnerabilityAssessmentRuleBaselinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseVulnerabilityAssessmentRuleBaselines client: %+v", err)
	}
	configureFunc(databaseVulnerabilityAssessmentRuleBaselinesClient.Client)

	databaseVulnerabilityAssessmentScansClient, err := databasevulnerabilityassessmentscans.NewDatabaseVulnerabilityAssessmentScansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseVulnerabilityAssessmentScans client: %+v", err)
	}
	configureFunc(databaseVulnerabilityAssessmentScansClient.Client)

	databaseVulnerabilityAssessmentsClient, err := databasevulnerabilityassessments.NewDatabaseVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(databaseVulnerabilityAssessmentsClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	deletedServersClient, err := deletedservers.NewDeletedServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedServers client: %+v", err)
	}
	configureFunc(deletedServersClient.Client)

	elasticPoolOperationsClient, err := elasticpooloperations.NewElasticPoolOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticPoolOperations client: %+v", err)
	}
	configureFunc(elasticPoolOperationsClient.Client)

	elasticPoolsClient, err := elasticpools.NewElasticPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ElasticPools client: %+v", err)
	}
	configureFunc(elasticPoolsClient.Client)

	encryptionProtectorsClient, err := encryptionprotectors.NewEncryptionProtectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EncryptionProtectors client: %+v", err)
	}
	configureFunc(encryptionProtectorsClient.Client)

	failoverGroupsClient, err := failovergroups.NewFailoverGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FailoverGroups client: %+v", err)
	}
	configureFunc(failoverGroupsClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	instanceFailoverGroupsClient, err := instancefailovergroups.NewInstanceFailoverGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InstanceFailoverGroups client: %+v", err)
	}
	configureFunc(instanceFailoverGroupsClient.Client)

	instancePoolsClient, err := instancepools.NewInstancePoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InstancePools client: %+v", err)
	}
	configureFunc(instancePoolsClient.Client)

	jobAgentsClient, err := jobagents.NewJobAgentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobAgents client: %+v", err)
	}
	configureFunc(jobAgentsClient.Client)

	jobCredentialsClient, err := jobcredentials.NewJobCredentialsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobCredentials client: %+v", err)
	}
	configureFunc(jobCredentialsClient.Client)

	jobExecutionsClient, err := jobexecutions.NewJobExecutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobExecutions client: %+v", err)
	}
	configureFunc(jobExecutionsClient.Client)

	jobStepExecutionsClient, err := jobstepexecutions.NewJobStepExecutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobStepExecutions client: %+v", err)
	}
	configureFunc(jobStepExecutionsClient.Client)

	jobStepsClient, err := jobsteps.NewJobStepsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobSteps client: %+v", err)
	}
	configureFunc(jobStepsClient.Client)

	jobTargetExecutionsClient, err := jobtargetexecutions.NewJobTargetExecutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobTargetExecutions client: %+v", err)
	}
	configureFunc(jobTargetExecutionsClient.Client)

	jobTargetGroupsClient, err := jobtargetgroups.NewJobTargetGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobTargetGroups client: %+v", err)
	}
	configureFunc(jobTargetGroupsClient.Client)

	jobVersionsClient, err := jobversions.NewJobVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building JobVersions client: %+v", err)
	}
	configureFunc(jobVersionsClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	ledgerDigestUploadsClient, err := ledgerdigestuploads.NewLedgerDigestUploadsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LedgerDigestUploads client: %+v", err)
	}
	configureFunc(ledgerDigestUploadsClient.Client)

	locationCapabilitiesClient, err := locationcapabilities.NewLocationCapabilitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocationCapabilities client: %+v", err)
	}
	configureFunc(locationCapabilitiesClient.Client)

	longTermRetentionBackupsClient, err := longtermretentionbackups.NewLongTermRetentionBackupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LongTermRetentionBackups client: %+v", err)
	}
	configureFunc(longTermRetentionBackupsClient.Client)

	longTermRetentionManagedInstanceBackupsClient, err := longtermretentionmanagedinstancebackups.NewLongTermRetentionManagedInstanceBackupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LongTermRetentionManagedInstanceBackups client: %+v", err)
	}
	configureFunc(longTermRetentionManagedInstanceBackupsClient.Client)

	longTermRetentionPoliciesClient, err := longtermretentionpolicies.NewLongTermRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LongTermRetentionPolicies client: %+v", err)
	}
	configureFunc(longTermRetentionPoliciesClient.Client)

	maintenanceWindowOptionsClient, err := maintenancewindowoptions.NewMaintenanceWindowOptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceWindowOptions client: %+v", err)
	}
	configureFunc(maintenanceWindowOptionsClient.Client)

	maintenanceWindowsClient, err := maintenancewindows.NewMaintenanceWindowsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaintenanceWindows client: %+v", err)
	}
	configureFunc(maintenanceWindowsClient.Client)

	managedBackupShortTermRetentionPoliciesClient, err := managedbackupshorttermretentionpolicies.NewManagedBackupShortTermRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedBackupShortTermRetentionPolicies client: %+v", err)
	}
	configureFunc(managedBackupShortTermRetentionPoliciesClient.Client)

	managedDatabaseColumnsClient, err := manageddatabasecolumns.NewManagedDatabaseColumnsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseColumns client: %+v", err)
	}
	configureFunc(managedDatabaseColumnsClient.Client)

	managedDatabaseQueriesClient, err := manageddatabasequeries.NewManagedDatabaseQueriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseQueries client: %+v", err)
	}
	configureFunc(managedDatabaseQueriesClient.Client)

	managedDatabaseRestoreDetailsClient, err := manageddatabaserestoredetails.NewManagedDatabaseRestoreDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseRestoreDetails client: %+v", err)
	}
	configureFunc(managedDatabaseRestoreDetailsClient.Client)

	managedDatabaseSchemasClient, err := manageddatabaseschemas.NewManagedDatabaseSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSchemas client: %+v", err)
	}
	configureFunc(managedDatabaseSchemasClient.Client)

	managedDatabaseSecurityAlertPoliciesClient, err := manageddatabasesecurityalertpolicies.NewManagedDatabaseSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(managedDatabaseSecurityAlertPoliciesClient.Client)

	managedDatabaseSecurityEventsClient, err := manageddatabasesecurityevents.NewManagedDatabaseSecurityEventsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSecurityEvents client: %+v", err)
	}
	configureFunc(managedDatabaseSecurityEventsClient.Client)

	managedDatabaseSensitivityLabelsClient, err := manageddatabasesensitivitylabels.NewManagedDatabaseSensitivityLabelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseSensitivityLabels client: %+v", err)
	}
	configureFunc(managedDatabaseSensitivityLabelsClient.Client)

	managedDatabaseTablesClient, err := manageddatabasetables.NewManagedDatabaseTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseTables client: %+v", err)
	}
	configureFunc(managedDatabaseTablesClient.Client)

	managedDatabaseTransparentDataEncryptionClient, err := manageddatabasetransparentdataencryption.NewManagedDatabaseTransparentDataEncryptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseTransparentDataEncryption client: %+v", err)
	}
	configureFunc(managedDatabaseTransparentDataEncryptionClient.Client)

	managedDatabaseVulnerabilityAssessmentRuleBaselinesClient, err := manageddatabasevulnerabilityassessmentrulebaselines.NewManagedDatabaseVulnerabilityAssessmentRuleBaselinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseVulnerabilityAssessmentRuleBaselines client: %+v", err)
	}
	configureFunc(managedDatabaseVulnerabilityAssessmentRuleBaselinesClient.Client)

	managedDatabaseVulnerabilityAssessmentScansClient, err := manageddatabasevulnerabilityassessmentscans.NewManagedDatabaseVulnerabilityAssessmentScansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseVulnerabilityAssessmentScans client: %+v", err)
	}
	configureFunc(managedDatabaseVulnerabilityAssessmentScansClient.Client)

	managedDatabaseVulnerabilityAssessmentsClient, err := manageddatabasevulnerabilityassessments.NewManagedDatabaseVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabaseVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(managedDatabaseVulnerabilityAssessmentsClient.Client)

	managedDatabasesClient, err := manageddatabases.NewManagedDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedDatabases client: %+v", err)
	}
	configureFunc(managedDatabasesClient.Client)

	managedInstanceAdministratorsClient, err := managedinstanceadministrators.NewManagedInstanceAdministratorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceAdministrators client: %+v", err)
	}
	configureFunc(managedInstanceAdministratorsClient.Client)

	managedInstanceAzureADOnlyAuthenticationsClient, err := managedinstanceazureadonlyauthentications.NewManagedInstanceAzureADOnlyAuthenticationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceAzureADOnlyAuthentications client: %+v", err)
	}
	configureFunc(managedInstanceAzureADOnlyAuthenticationsClient.Client)

	managedInstanceEncryptionProtectorsClient, err := managedinstanceencryptionprotectors.NewManagedInstanceEncryptionProtectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceEncryptionProtectors client: %+v", err)
	}
	configureFunc(managedInstanceEncryptionProtectorsClient.Client)

	managedInstanceKeysClient, err := managedinstancekeys.NewManagedInstanceKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceKeys client: %+v", err)
	}
	configureFunc(managedInstanceKeysClient.Client)

	managedInstanceLongTermRetentionPoliciesClient, err := managedinstancelongtermretentionpolicies.NewManagedInstanceLongTermRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceLongTermRetentionPolicies client: %+v", err)
	}
	configureFunc(managedInstanceLongTermRetentionPoliciesClient.Client)

	managedInstanceOperationsClient, err := managedinstanceoperations.NewManagedInstanceOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceOperations client: %+v", err)
	}
	configureFunc(managedInstanceOperationsClient.Client)

	managedInstancePrivateEndpointConnectionsClient, err := managedinstanceprivateendpointconnections.NewManagedInstancePrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstancePrivateEndpointConnections client: %+v", err)
	}
	configureFunc(managedInstancePrivateEndpointConnectionsClient.Client)

	managedInstancePrivateLinkResourcesClient, err := managedinstanceprivatelinkresources.NewManagedInstancePrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstancePrivateLinkResources client: %+v", err)
	}
	configureFunc(managedInstancePrivateLinkResourcesClient.Client)

	managedInstanceTdeCertificatesClient, err := managedinstancetdecertificates.NewManagedInstanceTdeCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceTdeCertificates client: %+v", err)
	}
	configureFunc(managedInstanceTdeCertificatesClient.Client)

	managedInstanceVulnerabilityAssessmentsClient, err := managedinstancevulnerabilityassessments.NewManagedInstanceVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstanceVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(managedInstanceVulnerabilityAssessmentsClient.Client)

	managedInstancesClient, err := managedinstances.NewManagedInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedInstances client: %+v", err)
	}
	configureFunc(managedInstancesClient.Client)

	managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient, err := managedrestorabledroppeddatabasebackupshorttermretentionpolicies.NewManagedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies client: %+v", err)
	}
	configureFunc(managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient.Client)

	managedServerSecurityAlertPoliciesClient, err := managedserversecurityalertpolicies.NewManagedServerSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(managedServerSecurityAlertPoliciesClient.Client)

	outboundFirewallRulesClient, err := outboundfirewallrules.NewOutboundFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OutboundFirewallRules client: %+v", err)
	}
	configureFunc(outboundFirewallRulesClient.Client)

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

	recoverableManagedDatabasesClient, err := recoverablemanageddatabases.NewRecoverableManagedDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecoverableManagedDatabases client: %+v", err)
	}
	configureFunc(recoverableManagedDatabasesClient.Client)

	replicationLinksClient, err := replicationlinks.NewReplicationLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReplicationLinks client: %+v", err)
	}
	configureFunc(replicationLinksClient.Client)

	restorableDroppedDatabasesClient, err := restorabledroppeddatabases.NewRestorableDroppedDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorableDroppedDatabases client: %+v", err)
	}
	configureFunc(restorableDroppedDatabasesClient.Client)

	restorableDroppedManagedDatabasesClient, err := restorabledroppedmanageddatabases.NewRestorableDroppedManagedDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorableDroppedManagedDatabases client: %+v", err)
	}
	configureFunc(restorableDroppedManagedDatabasesClient.Client)

	restorePointsClient, err := restorepoints.NewRestorePointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorePoints client: %+v", err)
	}
	configureFunc(restorePointsClient.Client)

	sensitivityLabelsClient, err := sensitivitylabels.NewSensitivityLabelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SensitivityLabels client: %+v", err)
	}
	configureFunc(sensitivityLabelsClient.Client)

	serverAdvisorsClient, err := serveradvisors.NewServerAdvisorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerAdvisors client: %+v", err)
	}
	configureFunc(serverAdvisorsClient.Client)

	serverAutomaticTuningClient, err := serverautomatictuning.NewServerAutomaticTuningClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerAutomaticTuning client: %+v", err)
	}
	configureFunc(serverAutomaticTuningClient.Client)

	serverAzureADAdministratorsClient, err := serverazureadadministrators.NewServerAzureADAdministratorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerAzureADAdministrators client: %+v", err)
	}
	configureFunc(serverAzureADAdministratorsClient.Client)

	serverAzureADOnlyAuthenticationsClient, err := serverazureadonlyauthentications.NewServerAzureADOnlyAuthenticationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerAzureADOnlyAuthentications client: %+v", err)
	}
	configureFunc(serverAzureADOnlyAuthenticationsClient.Client)

	serverDevOpsAuditClient, err := serverdevopsaudit.NewServerDevOpsAuditClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerDevOpsAudit client: %+v", err)
	}
	configureFunc(serverDevOpsAuditClient.Client)

	serverDnsAliasesClient, err := serverdnsaliases.NewServerDnsAliasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerDnsAliases client: %+v", err)
	}
	configureFunc(serverDnsAliasesClient.Client)

	serverKeysClient, err := serverkeys.NewServerKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerKeys client: %+v", err)
	}
	configureFunc(serverKeysClient.Client)

	serverOperationsClient, err := serveroperations.NewServerOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerOperations client: %+v", err)
	}
	configureFunc(serverOperationsClient.Client)

	serverSecurityAlertPoliciesClient, err := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(serverSecurityAlertPoliciesClient.Client)

	serverTrustGroupsClient, err := servertrustgroups.NewServerTrustGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerTrustGroups client: %+v", err)
	}
	configureFunc(serverTrustGroupsClient.Client)

	serverVulnerabilityAssessmentsClient, err := servervulnerabilityassessments.NewServerVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(serverVulnerabilityAssessmentsClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	sqlAgentClient, err := sqlagent.NewSqlAgentClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlAgent client: %+v", err)
	}
	configureFunc(sqlAgentClient.Client)

	subscriptionUsagesClient, err := subscriptionusages.NewSubscriptionUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubscriptionUsages client: %+v", err)
	}
	configureFunc(subscriptionUsagesClient.Client)

	syncAgentsClient, err := syncagents.NewSyncAgentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SyncAgents client: %+v", err)
	}
	configureFunc(syncAgentsClient.Client)

	syncGroupsClient, err := syncgroups.NewSyncGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SyncGroups client: %+v", err)
	}
	configureFunc(syncGroupsClient.Client)

	syncMembersClient, err := syncmembers.NewSyncMembersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SyncMembers client: %+v", err)
	}
	configureFunc(syncMembersClient.Client)

	tdeCertificatesClient, err := tdecertificates.NewTdeCertificatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TdeCertificates client: %+v", err)
	}
	configureFunc(tdeCertificatesClient.Client)

	timeZonesClient, err := timezones.NewTimeZonesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TimeZones client: %+v", err)
	}
	configureFunc(timeZonesClient.Client)

	transparentDataEncryptionsClient, err := transparentdataencryptions.NewTransparentDataEncryptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TransparentDataEncryptions client: %+v", err)
	}
	configureFunc(transparentDataEncryptionsClient.Client)

	usagesClient, err := usages.NewUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Usages client: %+v", err)
	}
	configureFunc(usagesClient.Client)

	virtualClustersClient, err := virtualclusters.NewVirtualClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualClusters client: %+v", err)
	}
	configureFunc(virtualClustersClient.Client)

	virtualNetworkRulesClient, err := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkRules client: %+v", err)
	}
	configureFunc(virtualNetworkRulesClient.Client)

	workloadClassifiersClient, err := workloadclassifiers.NewWorkloadClassifiersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadClassifiers client: %+v", err)
	}
	configureFunc(workloadClassifiersClient.Client)

	workloadGroupsClient, err := workloadgroups.NewWorkloadGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkloadGroups client: %+v", err)
	}
	configureFunc(workloadGroupsClient.Client)

	return &Client{
		BackupShortTermRetentionPolicies:             backupShortTermRetentionPoliciesClient,
		BlobAuditing:                                 blobAuditingClient,
		DataWarehouseUserActivities:                  dataWarehouseUserActivitiesClient,
		DatabaseAdvisors:                             databaseAdvisorsClient,
		DatabaseAutomaticTuning:                      databaseAutomaticTuningClient,
		DatabaseColumns:                              databaseColumnsClient,
		DatabaseExtensions:                           databaseExtensionsClient,
		DatabaseOperations:                           databaseOperationsClient,
		DatabaseRecommendedActions:                   databaseRecommendedActionsClient,
		DatabaseSchemas:                              databaseSchemasClient,
		DatabaseSecurityAlertPolicies:                databaseSecurityAlertPoliciesClient,
		DatabaseTables:                               databaseTablesClient,
		DatabaseUsages:                               databaseUsagesClient,
		DatabaseVulnerabilityAssessmentRuleBaselines: databaseVulnerabilityAssessmentRuleBaselinesClient,
		DatabaseVulnerabilityAssessmentScans:         databaseVulnerabilityAssessmentScansClient,
		DatabaseVulnerabilityAssessments:             databaseVulnerabilityAssessmentsClient,
		Databases:                                    databasesClient,
		DeletedServers:                               deletedServersClient,
		ElasticPoolOperations:                        elasticPoolOperationsClient,
		ElasticPools:                                 elasticPoolsClient,
		EncryptionProtectors:                         encryptionProtectorsClient,
		FailoverGroups:                               failoverGroupsClient,
		FirewallRules:                                firewallRulesClient,
		InstanceFailoverGroups:                       instanceFailoverGroupsClient,
		InstancePools:                                instancePoolsClient,
		JobAgents:                                    jobAgentsClient,
		JobCredentials:                               jobCredentialsClient,
		JobExecutions:                                jobExecutionsClient,
		JobStepExecutions:                            jobStepExecutionsClient,
		JobSteps:                                     jobStepsClient,
		JobTargetExecutions:                          jobTargetExecutionsClient,
		JobTargetGroups:                              jobTargetGroupsClient,
		JobVersions:                                  jobVersionsClient,
		Jobs:                                         jobsClient,
		LedgerDigestUploads:                          ledgerDigestUploadsClient,
		LocationCapabilities:                         locationCapabilitiesClient,
		LongTermRetentionBackups:                     longTermRetentionBackupsClient,
		LongTermRetentionManagedInstanceBackups:      longTermRetentionManagedInstanceBackupsClient,
		LongTermRetentionPolicies:                    longTermRetentionPoliciesClient,
		MaintenanceWindowOptions:                     maintenanceWindowOptionsClient,
		MaintenanceWindows:                           maintenanceWindowsClient,
		ManagedBackupShortTermRetentionPolicies:      managedBackupShortTermRetentionPoliciesClient,
		ManagedDatabaseColumns:                       managedDatabaseColumnsClient,
		ManagedDatabaseQueries:                       managedDatabaseQueriesClient,
		ManagedDatabaseRestoreDetails:                managedDatabaseRestoreDetailsClient,
		ManagedDatabaseSchemas:                       managedDatabaseSchemasClient,
		ManagedDatabaseSecurityAlertPolicies:         managedDatabaseSecurityAlertPoliciesClient,
		ManagedDatabaseSecurityEvents:                managedDatabaseSecurityEventsClient,
		ManagedDatabaseSensitivityLabels:             managedDatabaseSensitivityLabelsClient,
		ManagedDatabaseTables:                        managedDatabaseTablesClient,
		ManagedDatabaseTransparentDataEncryption:     managedDatabaseTransparentDataEncryptionClient,
		ManagedDatabaseVulnerabilityAssessmentRuleBaselines:              managedDatabaseVulnerabilityAssessmentRuleBaselinesClient,
		ManagedDatabaseVulnerabilityAssessmentScans:                      managedDatabaseVulnerabilityAssessmentScansClient,
		ManagedDatabaseVulnerabilityAssessments:                          managedDatabaseVulnerabilityAssessmentsClient,
		ManagedDatabases:                                                 managedDatabasesClient,
		ManagedInstanceAdministrators:                                    managedInstanceAdministratorsClient,
		ManagedInstanceAzureADOnlyAuthentications:                        managedInstanceAzureADOnlyAuthenticationsClient,
		ManagedInstanceEncryptionProtectors:                              managedInstanceEncryptionProtectorsClient,
		ManagedInstanceKeys:                                              managedInstanceKeysClient,
		ManagedInstanceLongTermRetentionPolicies:                         managedInstanceLongTermRetentionPoliciesClient,
		ManagedInstanceOperations:                                        managedInstanceOperationsClient,
		ManagedInstancePrivateEndpointConnections:                        managedInstancePrivateEndpointConnectionsClient,
		ManagedInstancePrivateLinkResources:                              managedInstancePrivateLinkResourcesClient,
		ManagedInstanceTdeCertificates:                                   managedInstanceTdeCertificatesClient,
		ManagedInstanceVulnerabilityAssessments:                          managedInstanceVulnerabilityAssessmentsClient,
		ManagedInstances:                                                 managedInstancesClient,
		ManagedRestorableDroppedDatabaseBackupShortTermRetentionPolicies: managedRestorableDroppedDatabaseBackupShortTermRetentionPoliciesClient,
		ManagedServerSecurityAlertPolicies:                               managedServerSecurityAlertPoliciesClient,
		OutboundFirewallRules:                                            outboundFirewallRulesClient,
		PrivateEndpointConnections:                                       privateEndpointConnectionsClient,
		PrivateLinkResources:                                             privateLinkResourcesClient,
		RecoverableManagedDatabases:                                      recoverableManagedDatabasesClient,
		ReplicationLinks:                                                 replicationLinksClient,
		RestorableDroppedDatabases:                                       restorableDroppedDatabasesClient,
		RestorableDroppedManagedDatabases:                                restorableDroppedManagedDatabasesClient,
		RestorePoints:                                                    restorePointsClient,
		SensitivityLabels:                                                sensitivityLabelsClient,
		ServerAdvisors:                                                   serverAdvisorsClient,
		ServerAutomaticTuning:                                            serverAutomaticTuningClient,
		ServerAzureADAdministrators:                                      serverAzureADAdministratorsClient,
		ServerAzureADOnlyAuthentications:                                 serverAzureADOnlyAuthenticationsClient,
		ServerDevOpsAudit:                                                serverDevOpsAuditClient,
		ServerDnsAliases:                                                 serverDnsAliasesClient,
		ServerKeys:                                                       serverKeysClient,
		ServerOperations:                                                 serverOperationsClient,
		ServerSecurityAlertPolicies:                                      serverSecurityAlertPoliciesClient,
		ServerTrustGroups:                                                serverTrustGroupsClient,
		ServerVulnerabilityAssessments:                                   serverVulnerabilityAssessmentsClient,
		Servers:                                                          serversClient,
		SqlAgent:                                                         sqlAgentClient,
		SubscriptionUsages:                                               subscriptionUsagesClient,
		SyncAgents:                                                       syncAgentsClient,
		SyncGroups:                                                       syncGroupsClient,
		SyncMembers:                                                      syncMembersClient,
		TdeCertificates:                                                  tdeCertificatesClient,
		TimeZones:                                                        timeZonesClient,
		TransparentDataEncryptions:                                       transparentDataEncryptionsClient,
		Usages:                                                           usagesClient,
		VirtualClusters:                                                  virtualClustersClient,
		VirtualNetworkRules:                                              virtualNetworkRulesClient,
		WorkloadClassifiers:                                              workloadClassifiersClient,
		WorkloadGroups:                                                   workloadGroupsClient,
	}, nil
}
