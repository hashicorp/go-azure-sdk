package virtualendpoint

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpoint", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointClient: %+v", err)
	}

	return &VirtualEndpointClient{
		Client: client,
	}, nil
}
