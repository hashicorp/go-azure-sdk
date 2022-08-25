package v2020_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2020-01-01/serverkeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2020-01-01/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2020-01-01/serverstop"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mysql/2020-01-01/serverupgrade"
)

type Client struct {
	ServerKeys    *serverkeys.ServerKeysClient
	ServerStart   *serverstart.ServerStartClient
	ServerStop    *serverstop.ServerStopClient
	ServerUpgrade *serverupgrade.ServerUpgradeClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	serverKeysClient := serverkeys.NewServerKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&serverKeysClient.Client)

	serverStartClient := serverstart.NewServerStartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStartClient.Client)

	serverStopClient := serverstop.NewServerStopClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStopClient.Client)

	serverUpgradeClient := serverupgrade.NewServerUpgradeClientWithBaseURI(endpoint)
	configureAuthFunc(&serverUpgradeClient.Client)

	return Client{
		ServerKeys:    &serverKeysClient,
		ServerStart:   &serverStartClient,
		ServerStop:    &serverStopClient,
		ServerUpgrade: &serverUpgradeClient,
	}
}
