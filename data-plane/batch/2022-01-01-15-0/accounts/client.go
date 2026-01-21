package accounts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountsClient struct {
	Client *dataplane.Client
}

func NewAccountsClientUnconfigured() (*AccountsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "", "accounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccountsClient: %+v", err)
	}

	return &AccountsClient{
		Client: client,
	}, nil
}

func (c *AccountsClient) AccountsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func (c *AccountsClient) AccountsClientSetAdditionalEndpoint(endpoint string) {
	c.Client.AdditionalEndpoint = endpoint
}

func NewAccountsClientWithBaseURI(endpoint string, additionalEndpoint string) (*AccountsClient, error) {
	client, err := dataplane.NewClient(endpoint, additionalEndpoint, "accounts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AccountsClient: %+v", err)
	}

	return &AccountsClient{
		Client: client,
	}, nil
}
