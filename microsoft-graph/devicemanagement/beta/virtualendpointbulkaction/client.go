package virtualendpointbulkaction

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEndpointBulkActionClient struct {
	Client *msgraph.Client
}

func NewVirtualEndpointBulkActionClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEndpointBulkActionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualendpointbulkaction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEndpointBulkActionClient: %+v", err)
	}

	return &VirtualEndpointBulkActionClient{
		Client: client,
	}, nil
}
