package v2018_06_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

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

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
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

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building LogFiles client: %+v", err)
	}
	configureFunc(logFilesClient.Client)

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

	updateConfigurationsClient, err := updateconfigurations.NewUpdateConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building UpdateConfigurations client: %+v", err)
	}
	configureFunc(updateConfigurationsClient.Client)

	virtualNetworkRulesClient, err := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkRules client: %+v", err)
	}
	configureFunc(virtualNetworkRulesClient.Client)

	return &Client{
		CheckNameAvailability:        checkNameAvailabilityClient,
		Configurations:               configurationsClient,
		Databases:                    databasesClient,
		FirewallRules:                firewallRulesClient,
		LocationBasedPerformanceTier: locationBasedPerformanceTierClient,
		LogFiles:                     logFilesClient,
		RecoverableServers:           recoverableServersClient,
		Replicas:                     replicasClient,
		ServerBasedPerformanceTier:   serverBasedPerformanceTierClient,
		ServerRestart:                serverRestartClient,
		ServerSecurityAlertPolicies:  serverSecurityAlertPoliciesClient,
		Servers:                      serversClient,
		UpdateConfigurations:         updateConfigurationsClient,
		VirtualNetworkRules:          virtualNetworkRulesClient,
	}, nil
}
