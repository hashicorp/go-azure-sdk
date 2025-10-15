package dataflowgraph

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphClient struct {
	Client *resourcemanager.Client
}

func NewDataflowGraphClientWithBaseURI(sdkApi sdkEnv.Api) (*DataflowGraphClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dataflowgraph", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataflowGraphClient: %+v", err)
	}

	return &DataflowGraphClient{
		Client: client,
	}, nil
}
