package akriconnector

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorClient struct {
	Client *resourcemanager.Client
}

func NewAkriConnectorClientWithBaseURI(sdkApi sdkEnv.Api) (*AkriConnectorClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "akriconnector", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AkriConnectorClient: %+v", err)
	}

	return &AkriConnectorClient{
		Client: client,
	}, nil
}
