package listprivatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListPrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewListPrivateLinkResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*ListPrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "listprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ListPrivateLinkResourcesClient: %+v", err)
	}

	return &ListPrivateLinkResourcesClient{
		Client: client,
	}, nil
}
