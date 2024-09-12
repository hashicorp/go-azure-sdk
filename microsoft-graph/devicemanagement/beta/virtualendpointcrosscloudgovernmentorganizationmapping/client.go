package virtualendpointcrosscloudgovernmentorganizationmapping

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointCrossCloudGovernmentOrganizationMappingClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointCrossCloudGovernmentOrganizationMappingClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointCrossCloudGovernmentOrganizationMappingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointcrosscloudgovernmentorganizationmapping", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointCrossCloudGovernmentOrganizationMappingClient: %+v", err)
	}

	return &VirtualEndpointCrossCloudGovernmentOrganizationMappingClient{
		Client: client,
	}, nil
}
