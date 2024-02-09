package nodereports

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeReportsClient struct {
	Client *resourcemanager.Client
}

func NewNodeReportsClientWithBaseURI(sdkApi sdkEnv.Api) (*NodeReportsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "nodereports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NodeReportsClient: %+v", err)
	}

	return &NodeReportsClient{
		Client: client,
	}, nil
}
