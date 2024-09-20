package getprivatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetPrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewGetPrivateLinkResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*GetPrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "getprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GetPrivateLinkResourcesClient: %+v", err)
	}

	return &GetPrivateLinkResourcesClient{
		Client: client,
	}, nil
}
