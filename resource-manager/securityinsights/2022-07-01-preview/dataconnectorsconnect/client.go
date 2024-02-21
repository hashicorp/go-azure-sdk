package dataconnectorsconnect

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorsConnectClient struct {
	Client *resourcemanager.Client
}

func NewDataConnectorsConnectClientWithBaseURI(sdkApi sdkEnv.Api) (*DataConnectorsConnectClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dataconnectorsconnect", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataConnectorsConnectClient: %+v", err)
	}

	return &DataConnectorsConnectClient{
		Client: client,
	}, nil
}
