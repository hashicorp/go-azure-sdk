package virtualendpointserviceplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointServicePlanClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointServicePlanClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointServicePlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointserviceplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointServicePlanClient: %+v", err)
	}

	return &VirtualEndpointServicePlanClient{
		Client: client,
	}, nil
}
