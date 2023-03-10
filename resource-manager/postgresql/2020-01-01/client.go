package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2020-01-01/serverkeys"
)

type Client struct {
	ServerKeys *serverkeys.ServerKeysClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	serverKeysClient := serverkeys.NewServerKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&serverKeysClient.Client)

	return Client{
		ServerKeys: &serverKeysClient,
	}
}
