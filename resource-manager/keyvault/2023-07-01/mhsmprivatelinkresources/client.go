package mhsmprivatelinkresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MHSMPrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewMHSMPrivateLinkResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*MHSMPrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "mhsmprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MHSMPrivateLinkResourcesClient: %+v", err)
	}

	return &MHSMPrivateLinkResourcesClient{
		Client: client,
	}, nil
}
