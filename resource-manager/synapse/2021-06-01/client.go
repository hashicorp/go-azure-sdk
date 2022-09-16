package v2021_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/bigdatapools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/databasevulnerabilityassesmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/integrationruntime"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/integrationruntimes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/ipfirewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/keys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/libraries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/privatelinkhubprivatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/privatelinkhubs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/recoverablesqlpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/restorabledroppedsqlpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sensitivitylabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsbackup"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsblobauditing"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsconnectionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatabasevulnerabilityassesmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatamaskingpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsdatamaskingrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsgeobackuppolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsmaintenancewindowoptions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsmaintenancewindows"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsreplicationlinks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsrestorepoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssensitivitylabels"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssqlpoolcolumns"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssqlpoolschemas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssqlpooltables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssqlpooluseractivities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssqlpoolvulnerabilityassesmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolssqlpoolvulnerabilityassessmentscans"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolstables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolstransparentdataencryption"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsvulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsvulnerabilityassessmentscansexecute"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsvulnerabilityassessmentscansexport"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsworkloadclassifiers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsworkloadgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspaceazureadonlyauthentications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserver"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverblobauditing"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverdedicatedsqlminimaltlssettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverserverencryptionprotector"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserverservervulnerabilityassessments"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspacemanagedsqlserversqlusages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/workspaces"
)

type Client struct {
	BigDataPools                                            *bigdatapools.BigDataPoolsClient
	DatabaseVulnerabilityAssesmentRuleBaselines             *databasevulnerabilityassesmentrulebaselines.DatabaseVulnerabilityAssesmentRuleBaselinesClient
	IPFirewallRules                                         *ipfirewallrules.IPFirewallRulesClient
	IntegrationRuntime                                      *integrationruntime.IntegrationRuntimeClient
	IntegrationRuntimes                                     *integrationruntimes.IntegrationRuntimesClient
	Keys                                                    *keys.KeysClient
	Libraries                                               *libraries.LibrariesClient
	PrivateEndpointConnections                              *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkHubPrivateLinkResources                      *privatelinkhubprivatelinkresources.PrivateLinkHubPrivateLinkResourcesClient
	PrivateLinkHubs                                         *privatelinkhubs.PrivateLinkHubsClient
	PrivateLinkResources                                    *privatelinkresources.PrivateLinkResourcesClient
	RecoverableSqlPools                                     *recoverablesqlpools.RecoverableSqlPoolsClient
	RestorableDroppedSqlPools                               *restorabledroppedsqlpools.RestorableDroppedSqlPoolsClient
	SensitivityLabels                                       *sensitivitylabels.SensitivityLabelsClient
	SqlPools                                                *sqlpools.SqlPoolsClient
	SqlPoolsBackup                                          *sqlpoolsbackup.SqlPoolsBackupClient
	SqlPoolsBlobAuditing                                    *sqlpoolsblobauditing.SqlPoolsBlobAuditingClient
	SqlPoolsConnectionPolicies                              *sqlpoolsconnectionpolicies.SqlPoolsConnectionPoliciesClient
	SqlPoolsDataMaskingPolicies                             *sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient
	SqlPoolsDataMaskingRules                                *sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient
	SqlPoolsDatabaseVulnerabilityAssesmentRuleBaselines     *sqlpoolsdatabasevulnerabilityassesmentrulebaselines.SqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient
	SqlPoolsGeoBackupPolicies                               *sqlpoolsgeobackuppolicies.SqlPoolsGeoBackupPoliciesClient
	SqlPoolsMaintenanceWindowOptions                        *sqlpoolsmaintenancewindowoptions.SqlPoolsMaintenanceWindowOptionsClient
	SqlPoolsMaintenanceWindows                              *sqlpoolsmaintenancewindows.SqlPoolsMaintenanceWindowsClient
	SqlPoolsReplicationLinks                                *sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient
	SqlPoolsRestorePoints                                   *sqlpoolsrestorepoints.SqlPoolsRestorePointsClient
	SqlPoolsSchemas                                         *sqlpoolsschemas.SqlPoolsSchemasClient
	SqlPoolsSecurityAlertPolicies                           *sqlpoolssecurityalertpolicies.SqlPoolsSecurityAlertPoliciesClient
	SqlPoolsSensitivityLabels                               *sqlpoolssensitivitylabels.SqlPoolsSensitivityLabelsClient
	SqlPoolsSqlPoolColumns                                  *sqlpoolssqlpoolcolumns.SqlPoolsSqlPoolColumnsClient
	SqlPoolsSqlPoolSchemas                                  *sqlpoolssqlpoolschemas.SqlPoolsSqlPoolSchemasClient
	SqlPoolsSqlPoolTables                                   *sqlpoolssqlpooltables.SqlPoolsSqlPoolTablesClient
	SqlPoolsSqlPoolUserActivities                           *sqlpoolssqlpooluseractivities.SqlPoolsSqlPoolUserActivitiesClient
	SqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselines      *sqlpoolssqlpoolvulnerabilityassesmentrulebaselines.SqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient
	SqlPoolsSqlPoolVulnerabilityAssessmentScans             *sqlpoolssqlpoolvulnerabilityassessmentscans.SqlPoolsSqlPoolVulnerabilityAssessmentScansClient
	SqlPoolsTables                                          *sqlpoolstables.SqlPoolsTablesClient
	SqlPoolsTransparentDataEncryption                       *sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient
	SqlPoolsUsages                                          *sqlpoolsusages.SqlPoolsUsagesClient
	SqlPoolsVulnerabilityAssessmentScansExecute             *sqlpoolsvulnerabilityassessmentscansexecute.SqlPoolsVulnerabilityAssessmentScansExecuteClient
	SqlPoolsVulnerabilityAssessmentScansExport              *sqlpoolsvulnerabilityassessmentscansexport.SqlPoolsVulnerabilityAssessmentScansExportClient
	SqlPoolsVulnerabilityAssessments                        *sqlpoolsvulnerabilityassessments.SqlPoolsVulnerabilityAssessmentsClient
	SqlPoolsWorkloadClassifiers                             *sqlpoolsworkloadclassifiers.SqlPoolsWorkloadClassifiersClient
	SqlPoolsWorkloadGroups                                  *sqlpoolsworkloadgroups.SqlPoolsWorkloadGroupsClient
	WorkspaceAzureADOnlyAuthentications                     *workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient
	WorkspaceManagedSqlServer                               *workspacemanagedsqlserver.WorkspaceManagedSqlServerClient
	WorkspaceManagedSqlServerBlobAuditing                   *workspacemanagedsqlserverblobauditing.WorkspaceManagedSqlServerBlobAuditingClient
	WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings *workspacemanagedsqlserverdedicatedsqlminimaltlssettings.WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient
	WorkspaceManagedSqlServerSecurityAlertPolicies          *workspacemanagedsqlserversecurityalertpolicies.WorkspaceManagedSqlServerSecurityAlertPoliciesClient
	WorkspaceManagedSqlServerServerEncryptionProtector      *workspacemanagedsqlserverserverencryptionprotector.WorkspaceManagedSqlServerServerEncryptionProtectorClient
	WorkspaceManagedSqlServerServerVulnerabilityAssessments *workspacemanagedsqlserverservervulnerabilityassessments.WorkspaceManagedSqlServerServerVulnerabilityAssessmentsClient
	WorkspaceManagedSqlServerSqlUsages                      *workspacemanagedsqlserversqlusages.WorkspaceManagedSqlServerSqlUsagesClient
	Workspaces                                              *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	bigDataPoolsClient := bigdatapools.NewBigDataPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&bigDataPoolsClient.Client)

	databaseVulnerabilityAssesmentRuleBaselinesClient := databasevulnerabilityassesmentrulebaselines.NewDatabaseVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(endpoint)
	configureAuthFunc(&databaseVulnerabilityAssesmentRuleBaselinesClient.Client)

	iPFirewallRulesClient := ipfirewallrules.NewIPFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&iPFirewallRulesClient.Client)

	integrationRuntimeClient := integrationruntime.NewIntegrationRuntimeClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationRuntimeClient.Client)

	integrationRuntimesClient := integrationruntimes.NewIntegrationRuntimesClientWithBaseURI(endpoint)
	configureAuthFunc(&integrationRuntimesClient.Client)

	keysClient := keys.NewKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&keysClient.Client)

	librariesClient := libraries.NewLibrariesClientWithBaseURI(endpoint)
	configureAuthFunc(&librariesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkHubPrivateLinkResourcesClient := privatelinkhubprivatelinkresources.NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkHubPrivateLinkResourcesClient.Client)

	privateLinkHubsClient := privatelinkhubs.NewPrivateLinkHubsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkHubsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	recoverableSqlPoolsClient := recoverablesqlpools.NewRecoverableSqlPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&recoverableSqlPoolsClient.Client)

	restorableDroppedSqlPoolsClient := restorabledroppedsqlpools.NewRestorableDroppedSqlPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&restorableDroppedSqlPoolsClient.Client)

	sensitivityLabelsClient := sensitivitylabels.NewSensitivityLabelsClientWithBaseURI(endpoint)
	configureAuthFunc(&sensitivityLabelsClient.Client)

	sqlPoolsBackupClient := sqlpoolsbackup.NewSqlPoolsBackupClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsBackupClient.Client)

	sqlPoolsBlobAuditingClient := sqlpoolsblobauditing.NewSqlPoolsBlobAuditingClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsBlobAuditingClient.Client)

	sqlPoolsClient := sqlpools.NewSqlPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsClient.Client)

	sqlPoolsConnectionPoliciesClient := sqlpoolsconnectionpolicies.NewSqlPoolsConnectionPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsConnectionPoliciesClient.Client)

	sqlPoolsDataMaskingPoliciesClient := sqlpoolsdatamaskingpolicies.NewSqlPoolsDataMaskingPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsDataMaskingPoliciesClient.Client)

	sqlPoolsDataMaskingRulesClient := sqlpoolsdatamaskingrules.NewSqlPoolsDataMaskingRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsDataMaskingRulesClient.Client)

	sqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient := sqlpoolsdatabasevulnerabilityassesmentrulebaselines.NewSqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient.Client)

	sqlPoolsGeoBackupPoliciesClient := sqlpoolsgeobackuppolicies.NewSqlPoolsGeoBackupPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsGeoBackupPoliciesClient.Client)

	sqlPoolsMaintenanceWindowOptionsClient := sqlpoolsmaintenancewindowoptions.NewSqlPoolsMaintenanceWindowOptionsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsMaintenanceWindowOptionsClient.Client)

	sqlPoolsMaintenanceWindowsClient := sqlpoolsmaintenancewindows.NewSqlPoolsMaintenanceWindowsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsMaintenanceWindowsClient.Client)

	sqlPoolsReplicationLinksClient := sqlpoolsreplicationlinks.NewSqlPoolsReplicationLinksClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsReplicationLinksClient.Client)

	sqlPoolsRestorePointsClient := sqlpoolsrestorepoints.NewSqlPoolsRestorePointsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsRestorePointsClient.Client)

	sqlPoolsSchemasClient := sqlpoolsschemas.NewSqlPoolsSchemasClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSchemasClient.Client)

	sqlPoolsSecurityAlertPoliciesClient := sqlpoolssecurityalertpolicies.NewSqlPoolsSecurityAlertPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSecurityAlertPoliciesClient.Client)

	sqlPoolsSensitivityLabelsClient := sqlpoolssensitivitylabels.NewSqlPoolsSensitivityLabelsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSensitivityLabelsClient.Client)

	sqlPoolsSqlPoolColumnsClient := sqlpoolssqlpoolcolumns.NewSqlPoolsSqlPoolColumnsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSqlPoolColumnsClient.Client)

	sqlPoolsSqlPoolSchemasClient := sqlpoolssqlpoolschemas.NewSqlPoolsSqlPoolSchemasClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSqlPoolSchemasClient.Client)

	sqlPoolsSqlPoolTablesClient := sqlpoolssqlpooltables.NewSqlPoolsSqlPoolTablesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSqlPoolTablesClient.Client)

	sqlPoolsSqlPoolUserActivitiesClient := sqlpoolssqlpooluseractivities.NewSqlPoolsSqlPoolUserActivitiesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSqlPoolUserActivitiesClient.Client)

	sqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient := sqlpoolssqlpoolvulnerabilityassesmentrulebaselines.NewSqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient.Client)

	sqlPoolsSqlPoolVulnerabilityAssessmentScansClient := sqlpoolssqlpoolvulnerabilityassessmentscans.NewSqlPoolsSqlPoolVulnerabilityAssessmentScansClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsSqlPoolVulnerabilityAssessmentScansClient.Client)

	sqlPoolsTablesClient := sqlpoolstables.NewSqlPoolsTablesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsTablesClient.Client)

	sqlPoolsTransparentDataEncryptionClient := sqlpoolstransparentdataencryption.NewSqlPoolsTransparentDataEncryptionClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsTransparentDataEncryptionClient.Client)

	sqlPoolsUsagesClient := sqlpoolsusages.NewSqlPoolsUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsUsagesClient.Client)

	sqlPoolsVulnerabilityAssessmentScansExecuteClient := sqlpoolsvulnerabilityassessmentscansexecute.NewSqlPoolsVulnerabilityAssessmentScansExecuteClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsVulnerabilityAssessmentScansExecuteClient.Client)

	sqlPoolsVulnerabilityAssessmentScansExportClient := sqlpoolsvulnerabilityassessmentscansexport.NewSqlPoolsVulnerabilityAssessmentScansExportClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsVulnerabilityAssessmentScansExportClient.Client)

	sqlPoolsVulnerabilityAssessmentsClient := sqlpoolsvulnerabilityassessments.NewSqlPoolsVulnerabilityAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsVulnerabilityAssessmentsClient.Client)

	sqlPoolsWorkloadClassifiersClient := sqlpoolsworkloadclassifiers.NewSqlPoolsWorkloadClassifiersClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsWorkloadClassifiersClient.Client)

	sqlPoolsWorkloadGroupsClient := sqlpoolsworkloadgroups.NewSqlPoolsWorkloadGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlPoolsWorkloadGroupsClient.Client)

	workspaceAzureADOnlyAuthenticationsClient := workspaceazureadonlyauthentications.NewWorkspaceAzureADOnlyAuthenticationsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceAzureADOnlyAuthenticationsClient.Client)

	workspaceManagedSqlServerBlobAuditingClient := workspacemanagedsqlserverblobauditing.NewWorkspaceManagedSqlServerBlobAuditingClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerBlobAuditingClient.Client)

	workspaceManagedSqlServerClient := workspacemanagedsqlserver.NewWorkspaceManagedSqlServerClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerClient.Client)

	workspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.NewWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient.Client)

	workspaceManagedSqlServerSecurityAlertPoliciesClient := workspacemanagedsqlserversecurityalertpolicies.NewWorkspaceManagedSqlServerSecurityAlertPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerSecurityAlertPoliciesClient.Client)

	workspaceManagedSqlServerServerEncryptionProtectorClient := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceManagedSqlServerServerEncryptionProtectorClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerServerEncryptionProtectorClient.Client)

	workspaceManagedSqlServerServerVulnerabilityAssessmentsClient := workspacemanagedsqlserverservervulnerabilityassessments.NewWorkspaceManagedSqlServerServerVulnerabilityAssessmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerServerVulnerabilityAssessmentsClient.Client)

	workspaceManagedSqlServerSqlUsagesClient := workspacemanagedsqlserversqlusages.NewWorkspaceManagedSqlServerSqlUsagesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceManagedSqlServerSqlUsagesClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		BigDataPools: &bigDataPoolsClient,
		DatabaseVulnerabilityAssesmentRuleBaselines: &databaseVulnerabilityAssesmentRuleBaselinesClient,
		IPFirewallRules:                    &iPFirewallRulesClient,
		IntegrationRuntime:                 &integrationRuntimeClient,
		IntegrationRuntimes:                &integrationRuntimesClient,
		Keys:                               &keysClient,
		Libraries:                          &librariesClient,
		PrivateEndpointConnections:         &privateEndpointConnectionsClient,
		PrivateLinkHubPrivateLinkResources: &privateLinkHubPrivateLinkResourcesClient,
		PrivateLinkHubs:                    &privateLinkHubsClient,
		PrivateLinkResources:               &privateLinkResourcesClient,
		RecoverableSqlPools:                &recoverableSqlPoolsClient,
		RestorableDroppedSqlPools:          &restorableDroppedSqlPoolsClient,
		SensitivityLabels:                  &sensitivityLabelsClient,
		SqlPools:                           &sqlPoolsClient,
		SqlPoolsBackup:                     &sqlPoolsBackupClient,
		SqlPoolsBlobAuditing:               &sqlPoolsBlobAuditingClient,
		SqlPoolsConnectionPolicies:         &sqlPoolsConnectionPoliciesClient,
		SqlPoolsDataMaskingPolicies:        &sqlPoolsDataMaskingPoliciesClient,
		SqlPoolsDataMaskingRules:           &sqlPoolsDataMaskingRulesClient,
		SqlPoolsDatabaseVulnerabilityAssesmentRuleBaselines:     &sqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient,
		SqlPoolsGeoBackupPolicies:                               &sqlPoolsGeoBackupPoliciesClient,
		SqlPoolsMaintenanceWindowOptions:                        &sqlPoolsMaintenanceWindowOptionsClient,
		SqlPoolsMaintenanceWindows:                              &sqlPoolsMaintenanceWindowsClient,
		SqlPoolsReplicationLinks:                                &sqlPoolsReplicationLinksClient,
		SqlPoolsRestorePoints:                                   &sqlPoolsRestorePointsClient,
		SqlPoolsSchemas:                                         &sqlPoolsSchemasClient,
		SqlPoolsSecurityAlertPolicies:                           &sqlPoolsSecurityAlertPoliciesClient,
		SqlPoolsSensitivityLabels:                               &sqlPoolsSensitivityLabelsClient,
		SqlPoolsSqlPoolColumns:                                  &sqlPoolsSqlPoolColumnsClient,
		SqlPoolsSqlPoolSchemas:                                  &sqlPoolsSqlPoolSchemasClient,
		SqlPoolsSqlPoolTables:                                   &sqlPoolsSqlPoolTablesClient,
		SqlPoolsSqlPoolUserActivities:                           &sqlPoolsSqlPoolUserActivitiesClient,
		SqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselines:      &sqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient,
		SqlPoolsSqlPoolVulnerabilityAssessmentScans:             &sqlPoolsSqlPoolVulnerabilityAssessmentScansClient,
		SqlPoolsTables:                                          &sqlPoolsTablesClient,
		SqlPoolsTransparentDataEncryption:                       &sqlPoolsTransparentDataEncryptionClient,
		SqlPoolsUsages:                                          &sqlPoolsUsagesClient,
		SqlPoolsVulnerabilityAssessmentScansExecute:             &sqlPoolsVulnerabilityAssessmentScansExecuteClient,
		SqlPoolsVulnerabilityAssessmentScansExport:              &sqlPoolsVulnerabilityAssessmentScansExportClient,
		SqlPoolsVulnerabilityAssessments:                        &sqlPoolsVulnerabilityAssessmentsClient,
		SqlPoolsWorkloadClassifiers:                             &sqlPoolsWorkloadClassifiersClient,
		SqlPoolsWorkloadGroups:                                  &sqlPoolsWorkloadGroupsClient,
		WorkspaceAzureADOnlyAuthentications:                     &workspaceAzureADOnlyAuthenticationsClient,
		WorkspaceManagedSqlServer:                               &workspaceManagedSqlServerClient,
		WorkspaceManagedSqlServerBlobAuditing:                   &workspaceManagedSqlServerBlobAuditingClient,
		WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings: &workspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient,
		WorkspaceManagedSqlServerSecurityAlertPolicies:          &workspaceManagedSqlServerSecurityAlertPoliciesClient,
		WorkspaceManagedSqlServerServerEncryptionProtector:      &workspaceManagedSqlServerServerEncryptionProtectorClient,
		WorkspaceManagedSqlServerServerVulnerabilityAssessments: &workspaceManagedSqlServerServerVulnerabilityAssessmentsClient,
		WorkspaceManagedSqlServerSqlUsages:                      &workspaceManagedSqlServerSqlUsagesClient,
		Workspaces:                                              &workspacesClient,
	}
}
