package privatelinkresourceoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkResourceOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkResourceOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateLinkResourceOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "privatelinkresourceoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkResourceOperationGroupClient: %+v", err)
	}

	return &PrivateLinkResourceOperationGroupClient{
		Client: client,
	}, nil
}
