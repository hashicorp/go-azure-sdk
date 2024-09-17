package dataconnectorsdisconnect

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsDisconnectClient struct {
	Client *resourcemanager.Client
}

func NewDataConnectorsDisconnectClientWithBaseURI(sdkApi sdkEnv.Api) (*DataConnectorsDisconnectClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dataconnectorsdisconnect", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataConnectorsDisconnectClient: %+v", err)
	}

	return &DataConnectorsDisconnectClient{
		Client: client,
	}, nil
}
