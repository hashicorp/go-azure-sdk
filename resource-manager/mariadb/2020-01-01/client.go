package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2020-01-01/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2020-01-01/serverstop"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	ServerStart *serverstart.ServerStartClient
	ServerStop  *serverstop.ServerStopClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	serverStartClient, err := serverstart.NewServerStartClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerStart client: %+v", err)
	}
	configureFunc(serverStartClient.Client)

	serverStopClient, err := serverstop.NewServerStopClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building ServerStop client: %+v", err)
	}
	configureFunc(serverStopClient.Client)

	return &Client{
		ServerStart: serverStartClient,
		ServerStop:  serverStopClient,
	}, nil
}
