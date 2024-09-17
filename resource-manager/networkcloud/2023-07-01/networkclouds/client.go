package networkclouds

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkcloudsClient struct {
	Client *resourcemanager.Client
}

func NewNetworkcloudsClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkcloudsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "networkclouds", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkcloudsClient: %+v", err)
	}

	return &NetworkcloudsClient{
		Client: client,
	}, nil
}
