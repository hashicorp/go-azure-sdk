package privatelinkhubprivatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkHubPrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkHubPrivateLinkResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateLinkHubPrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "privatelinkhubprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkHubPrivateLinkResourcesClient: %+v", err)
	}

	return &PrivateLinkHubPrivateLinkResourcesClient{
		Client: client,
	}, nil
}
