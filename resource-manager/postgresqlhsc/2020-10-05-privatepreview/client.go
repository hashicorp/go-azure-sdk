package v2020_10_05_privatepreview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/roles"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroupoperations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servergroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresqlhsc/2020-10-05-privatepreview/servers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Configurations        *configurations.ConfigurationsClient
	FirewallRules         *firewallrules.FirewallRulesClient
	Roles                 *roles.RolesClient
	ServerGroupOperations *servergroupoperations.ServerGroupOperationsClient
	ServerGroups          *servergroups.ServerGroupsClient
	Servers               *servers.ServersClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	rolesClient, err := roles.NewRolesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Roles client: %+v", err)
	}
	configureFunc(rolesClient.Client)

	serverGroupOperationsClient, err := servergroupoperations.NewServerGroupOperationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerGroupOperations client: %+v", err)
	}
	configureFunc(serverGroupOperationsClient.Client)

	serverGroupsClient, err := servergroups.NewServerGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerGroups client: %+v", err)
	}
	configureFunc(serverGroupsClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	return &Client{
		Configurations:        configurationsClient,
		FirewallRules:         firewallRulesClient,
		Roles:                 rolesClient,
		ServerGroupOperations: serverGroupOperationsClient,
		ServerGroups:          serverGroupsClient,
		Servers:               serversClient,
	}, nil
}
