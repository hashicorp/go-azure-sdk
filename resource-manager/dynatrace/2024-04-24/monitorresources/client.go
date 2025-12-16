package monitorresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorResourcesClient struct {
	Client *resourcemanager.Client
}

func NewMonitorResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*MonitorResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "monitorresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MonitorResourcesClient: %+v", err)
	}

	return &MonitorResourcesClient{
		Client: client,
	}, nil
}
