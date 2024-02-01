package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/bigdatapools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/databasevulnerabilityassesmentrulebaselines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/integrationruntime"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/integrationruntimes"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/ipfirewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/keys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/libraries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/operations"
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
	"github.com/hashicorp/go-azure-sdk/resource-manager/synapse/2021-06-01/sqlpoolsoperations"
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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	BigDataPools                                            *bigdatapools.BigDataPoolsClient
	DatabaseVulnerabilityAssesmentRuleBaselines             *databasevulnerabilityassesmentrulebaselines.DatabaseVulnerabilityAssesmentRuleBaselinesClient
	IPFirewallRules                                         *ipfirewallrules.IPFirewallRulesClient
	IntegrationRuntime                                      *integrationruntime.IntegrationRuntimeClient
	IntegrationRuntimes                                     *integrationruntimes.IntegrationRuntimesClient
	Keys                                                    *keys.KeysClient
	Libraries                                               *libraries.LibrariesClient
	Operations                                              *operations.OperationsClient
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
	SqlPoolsOperations                                      *sqlpoolsoperations.SqlPoolsOperationsClient
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	bigDataPoolsClient, err := bigdatapools.NewBigDataPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BigDataPools client: %+v", err)
	}
	configureFunc(bigDataPoolsClient.Client)

	databaseVulnerabilityAssesmentRuleBaselinesClient, err := databasevulnerabilityassesmentrulebaselines.NewDatabaseVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DatabaseVulnerabilityAssesmentRuleBaselines client: %+v", err)
	}
	configureFunc(databaseVulnerabilityAssesmentRuleBaselinesClient.Client)

	iPFirewallRulesClient, err := ipfirewallrules.NewIPFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IPFirewallRules client: %+v", err)
	}
	configureFunc(iPFirewallRulesClient.Client)

	integrationRuntimeClient, err := integrationruntime.NewIntegrationRuntimeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntime client: %+v", err)
	}
	configureFunc(integrationRuntimeClient.Client)

	integrationRuntimesClient, err := integrationruntimes.NewIntegrationRuntimesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IntegrationRuntimes client: %+v", err)
	}
	configureFunc(integrationRuntimesClient.Client)

	keysClient, err := keys.NewKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Keys client: %+v", err)
	}
	configureFunc(keysClient.Client)

	librariesClient, err := libraries.NewLibrariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Libraries client: %+v", err)
	}
	configureFunc(librariesClient.Client)

	operationsClient, err := operations.NewOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Operations client: %+v", err)
	}
	configureFunc(operationsClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkHubPrivateLinkResourcesClient, err := privatelinkhubprivatelinkresources.NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkHubPrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkHubPrivateLinkResourcesClient.Client)

	privateLinkHubsClient, err := privatelinkhubs.NewPrivateLinkHubsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkHubs client: %+v", err)
	}
	configureFunc(privateLinkHubsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	recoverableSqlPoolsClient, err := recoverablesqlpools.NewRecoverableSqlPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecoverableSqlPools client: %+v", err)
	}
	configureFunc(recoverableSqlPoolsClient.Client)

	restorableDroppedSqlPoolsClient, err := restorabledroppedsqlpools.NewRestorableDroppedSqlPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorableDroppedSqlPools client: %+v", err)
	}
	configureFunc(restorableDroppedSqlPoolsClient.Client)

	sensitivityLabelsClient, err := sensitivitylabels.NewSensitivityLabelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SensitivityLabels client: %+v", err)
	}
	configureFunc(sensitivityLabelsClient.Client)

	sqlPoolsBackupClient, err := sqlpoolsbackup.NewSqlPoolsBackupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsBackup client: %+v", err)
	}
	configureFunc(sqlPoolsBackupClient.Client)

	sqlPoolsBlobAuditingClient, err := sqlpoolsblobauditing.NewSqlPoolsBlobAuditingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsBlobAuditing client: %+v", err)
	}
	configureFunc(sqlPoolsBlobAuditingClient.Client)

	sqlPoolsClient, err := sqlpools.NewSqlPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPools client: %+v", err)
	}
	configureFunc(sqlPoolsClient.Client)

	sqlPoolsConnectionPoliciesClient, err := sqlpoolsconnectionpolicies.NewSqlPoolsConnectionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsConnectionPolicies client: %+v", err)
	}
	configureFunc(sqlPoolsConnectionPoliciesClient.Client)

	sqlPoolsDataMaskingPoliciesClient, err := sqlpoolsdatamaskingpolicies.NewSqlPoolsDataMaskingPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsDataMaskingPolicies client: %+v", err)
	}
	configureFunc(sqlPoolsDataMaskingPoliciesClient.Client)

	sqlPoolsDataMaskingRulesClient, err := sqlpoolsdatamaskingrules.NewSqlPoolsDataMaskingRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsDataMaskingRules client: %+v", err)
	}
	configureFunc(sqlPoolsDataMaskingRulesClient.Client)

	sqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient, err := sqlpoolsdatabasevulnerabilityassesmentrulebaselines.NewSqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsDatabaseVulnerabilityAssesmentRuleBaselines client: %+v", err)
	}
	configureFunc(sqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient.Client)

	sqlPoolsGeoBackupPoliciesClient, err := sqlpoolsgeobackuppolicies.NewSqlPoolsGeoBackupPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsGeoBackupPolicies client: %+v", err)
	}
	configureFunc(sqlPoolsGeoBackupPoliciesClient.Client)

	sqlPoolsMaintenanceWindowOptionsClient, err := sqlpoolsmaintenancewindowoptions.NewSqlPoolsMaintenanceWindowOptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsMaintenanceWindowOptions client: %+v", err)
	}
	configureFunc(sqlPoolsMaintenanceWindowOptionsClient.Client)

	sqlPoolsMaintenanceWindowsClient, err := sqlpoolsmaintenancewindows.NewSqlPoolsMaintenanceWindowsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsMaintenanceWindows client: %+v", err)
	}
	configureFunc(sqlPoolsMaintenanceWindowsClient.Client)

	sqlPoolsOperationsClient, err := sqlpoolsoperations.NewSqlPoolsOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsOperations client: %+v", err)
	}
	configureFunc(sqlPoolsOperationsClient.Client)

	sqlPoolsReplicationLinksClient, err := sqlpoolsreplicationlinks.NewSqlPoolsReplicationLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsReplicationLinks client: %+v", err)
	}
	configureFunc(sqlPoolsReplicationLinksClient.Client)

	sqlPoolsRestorePointsClient, err := sqlpoolsrestorepoints.NewSqlPoolsRestorePointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsRestorePoints client: %+v", err)
	}
	configureFunc(sqlPoolsRestorePointsClient.Client)

	sqlPoolsSchemasClient, err := sqlpoolsschemas.NewSqlPoolsSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSchemas client: %+v", err)
	}
	configureFunc(sqlPoolsSchemasClient.Client)

	sqlPoolsSecurityAlertPoliciesClient, err := sqlpoolssecurityalertpolicies.NewSqlPoolsSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(sqlPoolsSecurityAlertPoliciesClient.Client)

	sqlPoolsSensitivityLabelsClient, err := sqlpoolssensitivitylabels.NewSqlPoolsSensitivityLabelsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSensitivityLabels client: %+v", err)
	}
	configureFunc(sqlPoolsSensitivityLabelsClient.Client)

	sqlPoolsSqlPoolColumnsClient, err := sqlpoolssqlpoolcolumns.NewSqlPoolsSqlPoolColumnsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSqlPoolColumns client: %+v", err)
	}
	configureFunc(sqlPoolsSqlPoolColumnsClient.Client)

	sqlPoolsSqlPoolSchemasClient, err := sqlpoolssqlpoolschemas.NewSqlPoolsSqlPoolSchemasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSqlPoolSchemas client: %+v", err)
	}
	configureFunc(sqlPoolsSqlPoolSchemasClient.Client)

	sqlPoolsSqlPoolTablesClient, err := sqlpoolssqlpooltables.NewSqlPoolsSqlPoolTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSqlPoolTables client: %+v", err)
	}
	configureFunc(sqlPoolsSqlPoolTablesClient.Client)

	sqlPoolsSqlPoolUserActivitiesClient, err := sqlpoolssqlpooluseractivities.NewSqlPoolsSqlPoolUserActivitiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSqlPoolUserActivities client: %+v", err)
	}
	configureFunc(sqlPoolsSqlPoolUserActivitiesClient.Client)

	sqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient, err := sqlpoolssqlpoolvulnerabilityassesmentrulebaselines.NewSqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselines client: %+v", err)
	}
	configureFunc(sqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient.Client)

	sqlPoolsSqlPoolVulnerabilityAssessmentScansClient, err := sqlpoolssqlpoolvulnerabilityassessmentscans.NewSqlPoolsSqlPoolVulnerabilityAssessmentScansClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsSqlPoolVulnerabilityAssessmentScans client: %+v", err)
	}
	configureFunc(sqlPoolsSqlPoolVulnerabilityAssessmentScansClient.Client)

	sqlPoolsTablesClient, err := sqlpoolstables.NewSqlPoolsTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsTables client: %+v", err)
	}
	configureFunc(sqlPoolsTablesClient.Client)

	sqlPoolsTransparentDataEncryptionClient, err := sqlpoolstransparentdataencryption.NewSqlPoolsTransparentDataEncryptionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsTransparentDataEncryption client: %+v", err)
	}
	configureFunc(sqlPoolsTransparentDataEncryptionClient.Client)

	sqlPoolsUsagesClient, err := sqlpoolsusages.NewSqlPoolsUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsUsages client: %+v", err)
	}
	configureFunc(sqlPoolsUsagesClient.Client)

	sqlPoolsVulnerabilityAssessmentScansExecuteClient, err := sqlpoolsvulnerabilityassessmentscansexecute.NewSqlPoolsVulnerabilityAssessmentScansExecuteClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsVulnerabilityAssessmentScansExecute client: %+v", err)
	}
	configureFunc(sqlPoolsVulnerabilityAssessmentScansExecuteClient.Client)

	sqlPoolsVulnerabilityAssessmentScansExportClient, err := sqlpoolsvulnerabilityassessmentscansexport.NewSqlPoolsVulnerabilityAssessmentScansExportClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsVulnerabilityAssessmentScansExport client: %+v", err)
	}
	configureFunc(sqlPoolsVulnerabilityAssessmentScansExportClient.Client)

	sqlPoolsVulnerabilityAssessmentsClient, err := sqlpoolsvulnerabilityassessments.NewSqlPoolsVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(sqlPoolsVulnerabilityAssessmentsClient.Client)

	sqlPoolsWorkloadClassifiersClient, err := sqlpoolsworkloadclassifiers.NewSqlPoolsWorkloadClassifiersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsWorkloadClassifiers client: %+v", err)
	}
	configureFunc(sqlPoolsWorkloadClassifiersClient.Client)

	sqlPoolsWorkloadGroupsClient, err := sqlpoolsworkloadgroups.NewSqlPoolsWorkloadGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlPoolsWorkloadGroups client: %+v", err)
	}
	configureFunc(sqlPoolsWorkloadGroupsClient.Client)

	workspaceAzureADOnlyAuthenticationsClient, err := workspaceazureadonlyauthentications.NewWorkspaceAzureADOnlyAuthenticationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceAzureADOnlyAuthentications client: %+v", err)
	}
	configureFunc(workspaceAzureADOnlyAuthenticationsClient.Client)

	workspaceManagedSqlServerBlobAuditingClient, err := workspacemanagedsqlserverblobauditing.NewWorkspaceManagedSqlServerBlobAuditingClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServerBlobAuditing client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerBlobAuditingClient.Client)

	workspaceManagedSqlServerClient, err := workspacemanagedsqlserver.NewWorkspaceManagedSqlServerClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServer client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerClient.Client)

	workspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient, err := workspacemanagedsqlserverdedicatedsqlminimaltlssettings.NewWorkspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient.Client)

	workspaceManagedSqlServerSecurityAlertPoliciesClient, err := workspacemanagedsqlserversecurityalertpolicies.NewWorkspaceManagedSqlServerSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerSecurityAlertPoliciesClient.Client)

	workspaceManagedSqlServerServerEncryptionProtectorClient, err := workspacemanagedsqlserverserverencryptionprotector.NewWorkspaceManagedSqlServerServerEncryptionProtectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServerServerEncryptionProtector client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerServerEncryptionProtectorClient.Client)

	workspaceManagedSqlServerServerVulnerabilityAssessmentsClient, err := workspacemanagedsqlserverservervulnerabilityassessments.NewWorkspaceManagedSqlServerServerVulnerabilityAssessmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServerServerVulnerabilityAssessments client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerServerVulnerabilityAssessmentsClient.Client)

	workspaceManagedSqlServerSqlUsagesClient, err := workspacemanagedsqlserversqlusages.NewWorkspaceManagedSqlServerSqlUsagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceManagedSqlServerSqlUsages client: %+v", err)
	}
	configureFunc(workspaceManagedSqlServerSqlUsagesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		BigDataPools: bigDataPoolsClient,
		DatabaseVulnerabilityAssesmentRuleBaselines: databaseVulnerabilityAssesmentRuleBaselinesClient,
		IPFirewallRules:                    iPFirewallRulesClient,
		IntegrationRuntime:                 integrationRuntimeClient,
		IntegrationRuntimes:                integrationRuntimesClient,
		Keys:                               keysClient,
		Libraries:                          librariesClient,
		Operations:                         operationsClient,
		PrivateEndpointConnections:         privateEndpointConnectionsClient,
		PrivateLinkHubPrivateLinkResources: privateLinkHubPrivateLinkResourcesClient,
		PrivateLinkHubs:                    privateLinkHubsClient,
		PrivateLinkResources:               privateLinkResourcesClient,
		RecoverableSqlPools:                recoverableSqlPoolsClient,
		RestorableDroppedSqlPools:          restorableDroppedSqlPoolsClient,
		SensitivityLabels:                  sensitivityLabelsClient,
		SqlPools:                           sqlPoolsClient,
		SqlPoolsBackup:                     sqlPoolsBackupClient,
		SqlPoolsBlobAuditing:               sqlPoolsBlobAuditingClient,
		SqlPoolsConnectionPolicies:         sqlPoolsConnectionPoliciesClient,
		SqlPoolsDataMaskingPolicies:        sqlPoolsDataMaskingPoliciesClient,
		SqlPoolsDataMaskingRules:           sqlPoolsDataMaskingRulesClient,
		SqlPoolsDatabaseVulnerabilityAssesmentRuleBaselines:     sqlPoolsDatabaseVulnerabilityAssesmentRuleBaselinesClient,
		SqlPoolsGeoBackupPolicies:                               sqlPoolsGeoBackupPoliciesClient,
		SqlPoolsMaintenanceWindowOptions:                        sqlPoolsMaintenanceWindowOptionsClient,
		SqlPoolsMaintenanceWindows:                              sqlPoolsMaintenanceWindowsClient,
		SqlPoolsOperations:                                      sqlPoolsOperationsClient,
		SqlPoolsReplicationLinks:                                sqlPoolsReplicationLinksClient,
		SqlPoolsRestorePoints:                                   sqlPoolsRestorePointsClient,
		SqlPoolsSchemas:                                         sqlPoolsSchemasClient,
		SqlPoolsSecurityAlertPolicies:                           sqlPoolsSecurityAlertPoliciesClient,
		SqlPoolsSensitivityLabels:                               sqlPoolsSensitivityLabelsClient,
		SqlPoolsSqlPoolColumns:                                  sqlPoolsSqlPoolColumnsClient,
		SqlPoolsSqlPoolSchemas:                                  sqlPoolsSqlPoolSchemasClient,
		SqlPoolsSqlPoolTables:                                   sqlPoolsSqlPoolTablesClient,
		SqlPoolsSqlPoolUserActivities:                           sqlPoolsSqlPoolUserActivitiesClient,
		SqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselines:      sqlPoolsSqlPoolVulnerabilityAssesmentRuleBaselinesClient,
		SqlPoolsSqlPoolVulnerabilityAssessmentScans:             sqlPoolsSqlPoolVulnerabilityAssessmentScansClient,
		SqlPoolsTables:                                          sqlPoolsTablesClient,
		SqlPoolsTransparentDataEncryption:                       sqlPoolsTransparentDataEncryptionClient,
		SqlPoolsUsages:                                          sqlPoolsUsagesClient,
		SqlPoolsVulnerabilityAssessmentScansExecute:             sqlPoolsVulnerabilityAssessmentScansExecuteClient,
		SqlPoolsVulnerabilityAssessmentScansExport:              sqlPoolsVulnerabilityAssessmentScansExportClient,
		SqlPoolsVulnerabilityAssessments:                        sqlPoolsVulnerabilityAssessmentsClient,
		SqlPoolsWorkloadClassifiers:                             sqlPoolsWorkloadClassifiersClient,
		SqlPoolsWorkloadGroups:                                  sqlPoolsWorkloadGroupsClient,
		WorkspaceAzureADOnlyAuthentications:                     workspaceAzureADOnlyAuthenticationsClient,
		WorkspaceManagedSqlServer:                               workspaceManagedSqlServerClient,
		WorkspaceManagedSqlServerBlobAuditing:                   workspaceManagedSqlServerBlobAuditingClient,
		WorkspaceManagedSqlServerDedicatedSQLminimalTlsSettings: workspaceManagedSqlServerDedicatedSQLminimalTlsSettingsClient,
		WorkspaceManagedSqlServerSecurityAlertPolicies:          workspaceManagedSqlServerSecurityAlertPoliciesClient,
		WorkspaceManagedSqlServerServerEncryptionProtector:      workspaceManagedSqlServerServerEncryptionProtectorClient,
		WorkspaceManagedSqlServerServerVulnerabilityAssessments: workspaceManagedSqlServerServerVulnerabilityAssessmentsClient,
		WorkspaceManagedSqlServerSqlUsages:                      workspaceManagedSqlServerSqlUsagesClient,
		Workspaces:                                              workspacesClient,
	}, nil
}
