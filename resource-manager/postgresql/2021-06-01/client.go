package v2021_06_01

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/serverstop"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	CheckNameAvailability     *checknameavailability.CheckNameAvailabilityClient
	Configurations            *configurations.ConfigurationsClient
	Databases                 *databases.DatabasesClient
	FirewallRules             *firewallrules.FirewallRulesClient
	GetPrivateDnsZoneSuffix   *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities *locationbasedcapabilities.LocationBasedCapabilitiesClient
	ServerRestart             *serverrestart.ServerRestartClient
	ServerStart               *serverstart.ServerStartClient
	ServerStop                *serverstop.ServerStopClient
	Servers                   *servers.ServersClient
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

	getPrivateDnsZoneSuffixClient, err := getprivatednszonesuffix.NewGetPrivateDnsZoneSuffixClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for GetPrivateDnsZoneSuffix: %+v", err)
	}
	configureAuthFunc(getPrivateDnsZoneSuffixClient.Client)

	locationBasedCapabilitiesClient, err := locationbasedcapabilities.NewLocationBasedCapabilitiesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for LocationBasedCapabilities: %+v", err)
	}
	configureAuthFunc(locationBasedCapabilitiesClient.Client)

	serverRestartClient, err := serverrestart.NewServerRestartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerRestart: %+v", err)
	}
	configureAuthFunc(serverRestartClient.Client)

	serverStartClient, err := serverstart.NewServerStartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerStart: %+v", err)
	}
	configureAuthFunc(serverStartClient.Client)

	serverStopClient, err := serverstop.NewServerStopClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for ServerStop: %+v", err)
	}
	configureAuthFunc(serverStopClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Servers: %+v", err)
	}
	configureAuthFunc(serversClient.Client)

	return &Client{
		CheckNameAvailability:     checkNameAvailabilityClient,
		Configurations:            configurationsClient,
		Databases:                 databasesClient,
		FirewallRules:             firewallRulesClient,
		GetPrivateDnsZoneSuffix:   getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: locationBasedCapabilitiesClient,
		ServerRestart:             serverRestartClient,
		ServerStart:               serverStartClient,
		ServerStop:                serverStopClient,
		Servers:                   serversClient,
	}, nil
}
