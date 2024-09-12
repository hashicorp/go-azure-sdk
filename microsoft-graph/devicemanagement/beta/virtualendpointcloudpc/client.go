package virtualendpointcloudpc

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointCloudPCClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointCloudPCClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointCloudPCClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointcloudpc", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointCloudPCClient: %+v", err)
	}

	return &VirtualEndpointCloudPCClient{
		Client: client,
	}, nil
}
