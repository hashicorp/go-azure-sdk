package v2018_06_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/locationbasedperformancetier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/logfiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/recoverableservers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/serverbasedperformancetier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/serversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/updateconfigurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/virtualnetworkrules"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	CheckNameAvailability        *checknameavailability.CheckNameAvailabilityClient
	Configurations               *configurations.ConfigurationsClient
	Databases                    *databases.DatabasesClient
	FirewallRules                *firewallrules.FirewallRulesClient
	LocationBasedPerformanceTier *locationbasedperformancetier.LocationBasedPerformanceTierClient
	LogFiles                     *logfiles.LogFilesClient
	RecoverableServers           *recoverableservers.RecoverableServersClient
	Replicas                     *replicas.ReplicasClient
	ServerBasedPerformanceTier   *serverbasedperformancetier.ServerBasedPerformanceTierClient
	ServerRestart                *serverrestart.ServerRestartClient
	ServerSecurityAlertPolicies  *serversecurityalertpolicies.ServerSecurityAlertPoliciesClient
	Servers                      *servers.ServersClient
	UpdateConfigurations         *updateconfigurations.UpdateConfigurationsClient
	VirtualNetworkRules          *virtualnetworkrules.VirtualNetworkRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkNameAvailabilityClient := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkNameAvailabilityClient.Client)

	configurationsClient := configurations.NewConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&configurationsClient.Client)

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	firewallRulesClient := firewallrules.NewFirewallRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&firewallRulesClient.Client)

	locationBasedPerformanceTierClient := locationbasedperformancetier.NewLocationBasedPerformanceTierClientWithBaseURI(endpoint)
	configureAuthFunc(&locationBasedPerformanceTierClient.Client)

	logFilesClient := logfiles.NewLogFilesClientWithBaseURI(endpoint)
	configureAuthFunc(&logFilesClient.Client)

	recoverableServersClient := recoverableservers.NewRecoverableServersClientWithBaseURI(endpoint)
	configureAuthFunc(&recoverableServersClient.Client)

	replicasClient := replicas.NewReplicasClientWithBaseURI(endpoint)
	configureAuthFunc(&replicasClient.Client)

	serverBasedPerformanceTierClient := serverbasedperformancetier.NewServerBasedPerformanceTierClientWithBaseURI(endpoint)
	configureAuthFunc(&serverBasedPerformanceTierClient.Client)

	serverRestartClient := serverrestart.NewServerRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverRestartClient.Client)

	serverSecurityAlertPoliciesClient := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&serverSecurityAlertPoliciesClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	updateConfigurationsClient := updateconfigurations.NewUpdateConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&updateConfigurationsClient.Client)

	virtualNetworkRulesClient := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkRulesClient.Client)

	return Client{
		CheckNameAvailability:        &checkNameAvailabilityClient,
		Configurations:               &configurationsClient,
		Databases:                    &databasesClient,
		FirewallRules:                &firewallRulesClient,
		LocationBasedPerformanceTier: &locationBasedPerformanceTierClient,
		LogFiles:                     &logFilesClient,
		RecoverableServers:           &recoverableServersClient,
		Replicas:                     &replicasClient,
		ServerBasedPerformanceTier:   &serverBasedPerformanceTierClient,
		ServerRestart:                &serverRestartClient,
		ServerSecurityAlertPolicies:  &serverSecurityAlertPoliciesClient,
		Servers:                      &serversClient,
		UpdateConfigurations:         &updateConfigurationsClient,
		VirtualNetworkRules:          &virtualNetworkRulesClient,
	}
}
