package virtualendpointshareduseserviceplan

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointSharedUseServicePlanClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointSharedUseServicePlanClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointSharedUseServicePlanClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointshareduseserviceplan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointSharedUseServicePlanClient: %+v", err)
	}

	return &VirtualEndpointSharedUseServicePlanClient{
		Client: client,
	}, nil
}
