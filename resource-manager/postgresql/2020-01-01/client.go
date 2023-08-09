package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2020-01-01/serverkeys"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ServerKeys *serverkeys.ServerKeysClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	serverKeysClient, err := serverkeys.NewServerKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerKeys client: %+v", err)
	}
	configureFunc(serverKeysClient.Client)

	return &Client{
		ServerKeys: serverKeysClient,
	}, nil
}
