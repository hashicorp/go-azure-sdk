package v2024_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/signalr/2024-03-01/signalr"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	SignalR *signalr.SignalRClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	signalRClient, err := signalr.NewSignalRClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SignalR client: %+v", err)
	}
	configureFunc(signalRClient.Client)

	return &Client{
		SignalR: signalRClient,
	}, nil
}
