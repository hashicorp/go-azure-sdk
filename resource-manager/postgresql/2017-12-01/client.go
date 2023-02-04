package v2017_12_01

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for CheckNameAvailability: %+v", err)
	}
	configureAuthFunc(checkNameAvailabilityClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Configurations: %+v", err)
	}
	configureAuthFunc(configurationsClient.Client)

	configurationsUpdateClient, err := configurationsupdate.NewConfigurationsUpdateClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ConfigurationsUpdate: %+v", err)
	}
	configureAuthFunc(configurationsUpdateClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Databases: %+v", err)
	}
	configureAuthFunc(databasesClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for FirewallRules: %+v", err)
	}
	configureAuthFunc(firewallRulesClient.Client)

	locationBasedPerformanceTierClient, err := locationbasedperformancetier.NewLocationBasedPerformanceTierClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for LocationBasedPerformanceTier: %+v", err)
	}
	configureAuthFunc(locationBasedPerformanceTierClient.Client)

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for LogFiles: %+v", err)
	}
	configureAuthFunc(logFilesClient.Client)

	recoverableServersClient, err := recoverableservers.NewRecoverableServersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for RecoverableServers: %+v", err)
	}
	configureAuthFunc(recoverableServersClient.Client)

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Replicas: %+v", err)
	}
	configureAuthFunc(replicasClient.Client)

	serverAdministratorsClient, err := serveradministrators.NewServerAdministratorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerAdministrators: %+v", err)
	}
	configureAuthFunc(serverAdministratorsClient.Client)

	serverBasedPerformanceTierClient, err := serverbasedperformancetier.NewServerBasedPerformanceTierClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerBasedPerformanceTier: %+v", err)
	}
	configureAuthFunc(serverBasedPerformanceTierClient.Client)

	serverRestartClient, err := serverrestart.NewServerRestartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerRestart: %+v", err)
	}
	configureAuthFunc(serverRestartClient.Client)

	serverSecurityAlertPoliciesClient, err := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerSecurityAlertPolicies: %+v", err)
	}
	configureAuthFunc(serverSecurityAlertPoliciesClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Servers: %+v", err)
	}
	configureAuthFunc(serversClient.Client)

	virtualNetworkRulesClient, err := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for VirtualNetworkRules: %+v", err)
	}
	configureAuthFunc(virtualNetworkRulesClient.Client)

	return &Client{
		CheckNameAvailability:        checkNameAvailabilityClient,
		Configurations:               configurationsClient,
		ConfigurationsUpdate:         configurationsUpdateClient,
		Databases:                    databasesClient,
		FirewallRules:                firewallRulesClient,
		LocationBasedPerformanceTier: locationBasedPerformanceTierClient,
		LogFiles:                     logFilesClient,
		RecoverableServers:           recoverableServersClient,
		Replicas:                     replicasClient,
		ServerAdministrators:         serverAdministratorsClient,
		ServerBasedPerformanceTier:   serverBasedPerformanceTierClient,
		ServerRestart:                serverRestartClient,
		ServerSecurityAlertPolicies:  serverSecurityAlertPoliciesClient,
		Servers:                      serversClient,
		VirtualNetworkRules:          virtualNetworkRulesClient,
	}, nil
}
