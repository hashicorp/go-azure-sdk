package databoxedgedevices

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataBoxEdgeDevicesClient struct {
	Client *resourcemanager.Client
}

func NewDataBoxEdgeDevicesClientWithBaseURI(sdkApi sdkEnv.Api) (*DataBoxEdgeDevicesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "databoxedgedevices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataBoxEdgeDevicesClient: %+v", err)
	}

	return &DataBoxEdgeDevicesClient{
		Client: client,
	}, nil
}
