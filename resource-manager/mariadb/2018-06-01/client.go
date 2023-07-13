package v2018_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	advisorsClient, err := advisors.NewAdvisorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Advisors client: %+v", err)
	}
	configureFunc(advisorsClient.Client)

	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	configurationsUpdateClient, err := configurationsupdate.NewConfigurationsUpdateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ConfigurationsUpdate client: %+v", err)
	}
	configureFunc(configurationsUpdateClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	locationBasedPerformanceTierClient, err := locationbasedperformancetier.NewLocationBasedPerformanceTierClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LocationBasedPerformanceTier client: %+v", err)
	}
	configureFunc(locationBasedPerformanceTierClient.Client)

	locationBasedRecommendedActionSessionsResultClient, err := locationbasedrecommendedactionsessionsresult.NewLocationBasedRecommendedActionSessionsResultClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LocationBasedRecommendedActionSessionsResult client: %+v", err)
	}
	configureFunc(locationBasedRecommendedActionSessionsResultClient.Client)

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LogFiles client: %+v", err)
	}
	configureFunc(logFilesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	queryTextsClient, err := querytexts.NewQueryTextsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building QueryTexts client: %+v", err)
	}
	configureFunc(queryTextsClient.Client)

	recommendedActionSessionsClient, err := recommendedactionsessions.NewRecommendedActionSessionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RecommendedActionSessions client: %+v", err)
	}
	configureFunc(recommendedActionSessionsClient.Client)

	recommendedActionsClient, err := recommendedactions.NewRecommendedActionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RecommendedActions client: %+v", err)
	}
	configureFunc(recommendedActionsClient.Client)

	recoverableServersClient, err := recoverableservers.NewRecoverableServersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building RecoverableServers client: %+v", err)
	}
	configureFunc(recoverableServersClient.Client)

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Replicas client: %+v", err)
	}
	configureFunc(replicasClient.Client)

	resetQueryPerformanceInsightDataClient, err := resetqueryperformanceinsightdata.NewResetQueryPerformanceInsightDataClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ResetQueryPerformanceInsightData client: %+v", err)
	}
	configureFunc(resetQueryPerformanceInsightDataClient.Client)

	serverBasedPerformanceTierClient, err := serverbasedperformancetier.NewServerBasedPerformanceTierClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerBasedPerformanceTier client: %+v", err)
	}
	configureFunc(serverBasedPerformanceTierClient.Client)

	serverRestartClient, err := serverrestart.NewServerRestartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerRestart client: %+v", err)
	}
	configureFunc(serverRestartClient.Client)

	serverSecurityAlertPoliciesClient, err := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(serverSecurityAlertPoliciesClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	topQueryStatisticsClient, err := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building TopQueryStatistics client: %+v", err)
	}
	configureFunc(topQueryStatisticsClient.Client)

	virtualNetworkRulesClient, err := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkRules client: %+v", err)
	}
	configureFunc(virtualNetworkRulesClient.Client)

	waitStatisticsClient, err := waitstatistics.NewWaitStatisticsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building WaitStatistics client: %+v", err)
	}
	configureFunc(waitStatisticsClient.Client)

	return &Client{
		Advisors:                     advisorsClient,
		CheckNameAvailability:        checkNameAvailabilityClient,
		Configurations:               configurationsClient,
		ConfigurationsUpdate:         configurationsUpdateClient,
		Databases:                    databasesClient,
		FirewallRules:                firewallRulesClient,
		LocationBasedPerformanceTier: locationBasedPerformanceTierClient,
		LocationBasedRecommendedActionSessionsResult: locationBasedRecommendedActionSessionsResultClient,
		LogFiles:                         logFilesClient,
		PrivateEndpointConnections:       privateEndpointConnectionsClient,
		PrivateLinkResources:             privateLinkResourcesClient,
		QueryTexts:                       queryTextsClient,
		RecommendedActionSessions:        recommendedActionSessionsClient,
		RecommendedActions:               recommendedActionsClient,
		RecoverableServers:               recoverableServersClient,
		Replicas:                         replicasClient,
		ResetQueryPerformanceInsightData: resetQueryPerformanceInsightDataClient,
		ServerBasedPerformanceTier:       serverBasedPerformanceTierClient,
		ServerRestart:                    serverRestartClient,
		ServerSecurityAlertPolicies:      serverSecurityAlertPoliciesClient,
		Servers:                          serversClient,
		TopQueryStatistics:               topQueryStatisticsClient,
		VirtualNetworkRules:              virtualNetworkRulesClient,
		WaitStatistics:                   waitStatisticsClient,
	}, nil
}
