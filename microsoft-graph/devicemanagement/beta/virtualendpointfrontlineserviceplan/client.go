package virtualendpointfrontlineserviceplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointFrontLineServicePlanClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointFrontLineServicePlanClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointFrontLineServicePlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointfrontlineserviceplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointFrontLineServicePlanClient: %+v", err)
	}

	return &VirtualEndpointFrontLineServicePlanClient{
		Client: client,
	}, nil
}
