package v2021_02_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maps/2021-02-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maps/2021-02-01/creators"
)

type Client struct {
	Accounts *accounts.AccountsClient
	Creators *creators.CreatorsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accountsClient := accounts.NewAccountsClientWithBaseURI(endpoint)
	configureAuthFunc(&accountsClient.Client)

	creatorsClient := creators.NewCreatorsClientWithBaseURI(endpoint)
	configureAuthFunc(&creatorsClient.Client)

	return Client{
		Accounts: &accountsClient,
		Creators: &creatorsClient,
	}
}
