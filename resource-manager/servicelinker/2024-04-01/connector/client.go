package connector

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorClient struct {
	Client *resourcemanager.Client
}

func NewConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectorClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "connector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectorClient: %+v", err)
	}

	return &ConnectorClient{
		Client: client,
	}, nil
}
