package v2022_03_08_preview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/administrators"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/backups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/checknameavailability"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/configurations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/databases"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/firewallrules"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/getprivatednszonesuffix"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/locationbasedcapabilities"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/replicas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/serverrestart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/servers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2022-03-08-preview/serverstop"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Administrators            *administrators.AdministratorsClient
	Backups                   *backups.BackupsClient
	CheckNameAvailability     *checknameavailability.CheckNameAvailabilityClient
	Configurations            *configurations.ConfigurationsClient
	Databases                 *databases.DatabasesClient
	FirewallRules             *firewallrules.FirewallRulesClient
	GetPrivateDnsZoneSuffix   *getprivatednszonesuffix.GetPrivateDnsZoneSuffixClient
	LocationBasedCapabilities *locationbasedcapabilities.LocationBasedCapabilitiesClient
	Replicas                  *replicas.ReplicasClient
	ServerRestart             *serverrestart.ServerRestartClient
	ServerStart               *serverstart.ServerStartClient
	ServerStop                *serverstop.ServerStopClient
	Servers                   *servers.ServersClient
}

func NewClientWithBaseURI(api environments.Api, configureAuthFunc func(c *resourcemanager.Client)) (*Client, error) {
	administratorsClient, err := administrators.NewAdministratorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Administrators: %+v", err)
	}
	configureAuthFunc(administratorsClient.Client)

	backupsClient, err := backups.NewBackupsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Backups: %+v", err)
	}
	configureAuthFunc(backupsClient.Client)

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

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building meta client for Replicas: %+v", err)
	}
	configureAuthFunc(replicasClient.Client)

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
		Administrators:            administratorsClient,
		Backups:                   backupsClient,
		CheckNameAvailability:     checkNameAvailabilityClient,
		Configurations:            configurationsClient,
		Databases:                 databasesClient,
		FirewallRules:             firewallRulesClient,
		GetPrivateDnsZoneSuffix:   getPrivateDnsZoneSuffixClient,
		LocationBasedCapabilities: locationBasedCapabilitiesClient,
		Replicas:                  replicasClient,
		ServerRestart:             serverRestartClient,
		ServerStart:               serverStartClient,
		ServerStop:                serverStopClient,
		Servers:                   serversClient,
	}, nil
}
