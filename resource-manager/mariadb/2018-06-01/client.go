package v2018_06_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/advisors"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/configurationsupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/locationbasedperformancetier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/locationbasedrecommendedactionsessionsresult"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/logfiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/querytexts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/recommendedactions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/recommendedactionsessions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/recoverableservers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/resetqueryperformanceinsightdata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/serverbasedperformancetier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/serversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/topquerystatistics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/virtualnetworkrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01/waitstatistics"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	Advisors                                     *advisors.AdvisorsClient
	CheckNameAvailability                        *checknameavailability.CheckNameAvailabilityClient
	Configurations                               *configurations.ConfigurationsClient
	ConfigurationsUpdate                         *configurationsupdate.ConfigurationsUpdateClient
	Databases                                    *databases.DatabasesClient
	FirewallRules                                *firewallrules.FirewallRulesClient
	LocationBasedPerformanceTier                 *locationbasedperformancetier.LocationBasedPerformanceTierClient
	LocationBasedRecommendedActionSessionsResult *locationbasedrecommendedactionsessionsresult.LocationBasedRecommendedActionSessionsResultClient
	LogFiles                                     *logfiles.LogFilesClient
	PrivateEndpointConnections                   *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                         *privatelinkresources.PrivateLinkResourcesClient
	QueryTexts                                   *querytexts.QueryTextsClient
	RecommendedActionSessions                    *recommendedactionsessions.RecommendedActionSessionsClient
	RecommendedActions                           *recommendedactions.RecommendedActionsClient
	RecoverableServers                           *recoverableservers.RecoverableServersClient
	Replicas                                     *replicas.ReplicasClient
	ResetQueryPerformanceInsightData             *resetqueryperformanceinsightdata.ResetQueryPerformanceInsightDataClient
	ServerBasedPerformanceTier                   *serverbasedperformancetier.ServerBasedPerformanceTierClient
	ServerRestart                                *serverrestart.ServerRestartClient
	ServerSecurityAlertPolicies                  *serversecurityalertpolicies.ServerSecurityAlertPoliciesClient
	Servers                                      *servers.ServersClient
	TopQueryStatistics                           *topquerystatistics.TopQueryStatisticsClient
	VirtualNetworkRules                          *virtualnetworkrules.VirtualNetworkRulesClient
	WaitStatistics                               *waitstatistics.WaitStatisticsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	advisorsClient := advisors.NewAdvisorsClientWithBaseURI(endpoint)
	configureAuthFunc(&advisorsClient.Client)

	checkNameAvailabilityClient := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityClient.Client)

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	configurationsUpdateClient := configurationsupdate.NewConfigurationsUpdateClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsUpdateClient.Client)

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	locationBasedPerformanceTierClient := locationbasedperformancetier.NewLocationBasedPerformanceTierClientWithBaseURI(endpoint)
	configureAuthFunc(&locationBasedPerformanceTierClient.Client)

	locationBasedRecommendedActionSessionsResultClient := locationbasedrecommendedactionsessionsresult.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(endpoint)
	configureAuthFunc(&locationBasedRecommendedActionSessionsResultClient.Client)

	logFilesClient := logfiles.NewLogFilesClientWithBaseURI(endpoint)
	configureAuthFunc(&logFilesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	queryTextsClient := querytexts.NewQueryTextsClientWithBaseURI(endpoint)
	configureAuthFunc(&queryTextsClient.Client)

	recommendedActionSessionsClient := recommendedactionsessions.NewRecommendedActionSessionsClientWithBaseURI(endpoint)
	configureAuthFunc(&recommendedActionSessionsClient.Client)

	recommendedActionsClient := recommendedactions.NewRecommendedActionsClientWithBaseURI(endpoint)
	configureAuthFunc(&recommendedActionsClient.Client)

	recoverableServersClient := recoverableservers.NewRecoverableServersClientWithBaseURI(endpoint)
	configureAuthFunc(&recoverableServersClient.Client)

	replicasClient := replicas.NewReplicasClientWithBaseURI(endpoint)
	configureAuthFunc(&replicasClient.Client)

	resetQueryPerformanceInsightDataClient := resetqueryperformanceinsightdata.NewResetQueryPerformanceInsightDataClientWithBaseURI(endpoint)
	configureAuthFunc(&resetQueryPerformanceInsightDataClient.Client)

	serverBasedPerformanceTierClient := serverbasedperformancetier.NewServerBasedPerformanceTierClientWithBaseURI(endpoint)
	configureAuthFunc(&serverBasedPerformanceTierClient.Client)

	serverRestartClient := serverrestart.NewServerRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverRestartClient.Client)

	serverSecurityAlertPoliciesClient := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&serverSecurityAlertPoliciesClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	topQueryStatisticsClient := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&topQueryStatisticsClient.Client)

	virtualNetworkRulesClient := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkRulesClient.Client)

	waitStatisticsClient := waitstatistics.NewWaitStatisticsClientWithBaseURI(endpoint)
	configureAuthFunc(&waitStatisticsClient.Client)

	return Client{
		Advisors:                     &advisorsClient,
		CheckNameAvailability:        &checkNameAvailabilityClient,
		Configurations:               &configurationsClient,
		ConfigurationsUpdate:         &configurationsUpdateClient,
		Databases:                    &databasesClient,
		FirewallRules:                &firewallRulesClient,
		LocationBasedPerformanceTier: &locationBasedPerformanceTierClient,
		LocationBasedRecommendedActionSessionsResult: &locationBasedRecommendedActionSessionsResultClient,
		LogFiles:                         &logFilesClient,
		PrivateEndpointConnections:       &privateEndpointConnectionsClient,
		PrivateLinkResources:             &privateLinkResourcesClient,
		QueryTexts:                       &queryTextsClient,
		RecommendedActionSessions:        &recommendedActionSessionsClient,
		RecommendedActions:               &recommendedActionsClient,
		RecoverableServers:               &recoverableServersClient,
		Replicas:                         &replicasClient,
		ResetQueryPerformanceInsightData: &resetQueryPerformanceInsightDataClient,
		ServerBasedPerformanceTier:       &serverBasedPerformanceTierClient,
		ServerRestart:                    &serverRestartClient,
		ServerSecurityAlertPolicies:      &serverSecurityAlertPoliciesClient,
		Servers:                          &serversClient,
		TopQueryStatistics:               &topQueryStatisticsClient,
		VirtualNetworkRules:              &virtualNetworkRulesClient,
		WaitStatistics:                   &waitStatisticsClient,
	}
}
