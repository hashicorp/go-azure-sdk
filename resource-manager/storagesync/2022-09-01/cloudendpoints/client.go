package cloudendpoints

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointsClient struct {
	Client *resourcemanager.Client
}

func NewCloudEndpointsClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudEndpointsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "cloudendpoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudEndpointsClient: %+v", err)
	}

	return &CloudEndpointsClient{
		Client: client,
	}, nil
}
