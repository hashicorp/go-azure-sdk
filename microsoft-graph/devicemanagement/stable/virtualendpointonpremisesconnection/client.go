package virtualendpointonpremisesconnection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointOnPremisesConnectionClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointOnPremisesConnectionClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointOnPremisesConnectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointonpremisesconnection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointOnPremisesConnectionClient: %+v", err)
	}

	return &VirtualEndpointOnPremisesConnectionClient{
		Client: client,
	}, nil
}
