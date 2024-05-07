package routinginfo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoutingInfoClient struct {
	Client *resourcemanager.Client
}

func NewRoutingInfoClientWithBaseURI(sdkApi sdkEnv.Api) (*RoutingInfoClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "routinginfo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RoutingInfoClient: %+v", err)
	}

	return &RoutingInfoClient{
		Client: client,
	}, nil
}
