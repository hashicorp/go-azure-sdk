package v2021_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/maps/2021-02-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/maps/2021-02-01/creators"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Accounts *accounts.AccountsClient
	Creators *creators.CreatorsClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountsClient, err := accounts.NewAccountsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	creatorsClient, err := creators.NewCreatorsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Creators client: %+v", err)
	}
	configureFunc(creatorsClient.Client)

	return &Client{
		Accounts: accountsClient,
		Creators: creatorsClient,
	}, nil
}
