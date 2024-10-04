package v2024_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2024-01-01/accesstoken"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2024-01-01/accounts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/videoindexer/2024-01-01/nameavailability"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessToken      *accesstoken.AccessTokenClient
	Accounts         *accounts.AccountsClient
	NameAvailability *nameavailability.NameAvailabilityClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accessTokenClient, err := accesstoken.NewAccessTokenClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessToken client: %+v", err)
	}
	configureFunc(accessTokenClient.Client)

	accountsClient, err := accounts.NewAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Accounts client: %+v", err)
	}
	configureFunc(accountsClient.Client)

	nameAvailabilityClient, err := nameavailability.NewNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NameAvailability client: %+v", err)
	}
	configureFunc(nameAvailabilityClient.Client)

	return &Client{
		AccessToken:      accessTokenClient,
		Accounts:         accountsClient,
		NameAvailability: nameAvailabilityClient,
	}, nil
}
