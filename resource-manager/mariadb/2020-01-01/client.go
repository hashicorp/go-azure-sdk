// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2020_01_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2020-01-01/serverstart"
	"github.com/hashicorp/go-azure-sdk/resource-manager/mariadb/2020-01-01/serverstop"
)

type Client struct {
	ServerStart *serverstart.ServerStartClient
	ServerStop  *serverstop.ServerStopClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	serverStartClient := serverstart.NewServerStartClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStartClient.Client)

	serverStopClient := serverstop.NewServerStopClientWithBaseURI(endpoint)
	configureAuthFunc(&serverStopClient.Client)

	return Client{
		ServerStart: &serverStartClient,
		ServerStop:  &serverStopClient,
	}
}
