package virtualendpointsupportedregion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointSupportedRegionClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointSupportedRegionClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointSupportedRegionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointsupportedregion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointSupportedRegionClient: %+v", err)
	}

	return &VirtualEndpointSupportedRegionClient{
		Client: client,
	}, nil
}
