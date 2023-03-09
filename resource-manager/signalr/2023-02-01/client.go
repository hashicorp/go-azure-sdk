package v2023_02_01

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2023-02-01/signalr"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	SignalR *signalr.SignalRClient
}

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	signalRClient, err := signalr.NewSignalRClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building SignalR client: %+v", err)
	}
	configureFunc(signalRClient.Client)

	return &Client{
		SignalR: signalRClient,
	}, nil
}
