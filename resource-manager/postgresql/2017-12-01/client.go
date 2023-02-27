// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2017_12_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/configurationsupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/locationbasedperformancetier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/logfiles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/recoverableservers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/serveradministrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/serverbasedperformancetier"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/serversecurityalertpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2017-12-01/virtualnetworkrules"
)

type Client struct {
	CheckNameAvailability        *checknameavailability.CheckNameAvailabilityClient
	Configurations               *configurations.ConfigurationsClient
	ConfigurationsUpdate         *configurationsupdate.ConfigurationsUpdateClient
	Databases                    *databases.DatabasesClient
	FirewallRules                *firewallrules.FirewallRulesClient
	LocationBasedPerformanceTier *locationbasedperformancetier.LocationBasedPerformanceTierClient
	LogFiles                     *logfiles.LogFilesClient
	RecoverableServers           *recoverableservers.RecoverableServersClient
	Replicas                     *replicas.ReplicasClient
	ServerAdministrators         *serveradministrators.ServerAdministratorsClient
	ServerBasedPerformanceTier   *serverbasedperformancetier.ServerBasedPerformanceTierClient
	ServerRestart                *serverrestart.ServerRestartClient
	ServerSecurityAlertPolicies  *serversecurityalertpolicies.ServerSecurityAlertPoliciesClient
	Servers                      *servers.ServersClient
	VirtualNetworkRules          *virtualnetworkrules.VirtualNetworkRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

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

	logFilesClient := logfiles.NewLogFilesClientWithBaseURI(endpoint)
	configureAuthFunc(&logFilesClient.Client)

	recoverableServersClient := recoverableservers.NewRecoverableServersClientWithBaseURI(endpoint)
	configureAuthFunc(&recoverableServersClient.Client)

	replicasClient := replicas.NewReplicasClientWithBaseURI(endpoint)
	configureAuthFunc(&replicasClient.Client)

	serverAdministratorsClient := serveradministrators.NewServerAdministratorsClientWithBaseURI(endpoint)
	configureAuthFunc(&serverAdministratorsClient.Client)

	serverBasedPerformanceTierClient := serverbasedperformancetier.NewServerBasedPerformanceTierClientWithBaseURI(endpoint)
	configureAuthFunc(&serverBasedPerformanceTierClient.Client)

	serverRestartClient := serverrestart.NewServerRestartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverRestartClient.Client)

	serverSecurityAlertPoliciesClient := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&serverSecurityAlertPoliciesClient.Client)

	serversClient := servers.NewServersClientWithBaseURI(endpoint)
	configureAuthFunc(&serversClient.Client)

	virtualNetworkRulesClient := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworkRulesClient.Client)

	return Client{
		CheckNameAvailability:        &checkNameAvailabilityClient,
		Configurations:               &configurationsClient,
		ConfigurationsUpdate:         &configurationsUpdateClient,
		Databases:                    &databasesClient,
		FirewallRules:                &firewallRulesClient,
		LocationBasedPerformanceTier: &locationBasedPerformanceTierClient,
		LogFiles:                     &logFilesClient,
		RecoverableServers:           &recoverableServersClient,
		Replicas:                     &replicasClient,
		ServerAdministrators:         &serverAdministratorsClient,
		ServerBasedPerformanceTier:   &serverBasedPerformanceTierClient,
		ServerRestart:                &serverRestartClient,
		ServerSecurityAlertPolicies:  &serverSecurityAlertPoliciesClient,
		Servers:                      &serversClient,
		VirtualNetworkRules:          &virtualNetworkRulesClient,
	}
}
