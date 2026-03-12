package swiftvirtualnetworks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwiftVirtualNetworksClient struct {
	Client *resourcemanager.Client
}

func NewSwiftVirtualNetworksClientWithBaseURI(sdkApi sdkEnv.Api) (*SwiftVirtualNetworksClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "swiftvirtualnetworks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SwiftVirtualNetworksClient: %+v", err)
	}

	return &SwiftVirtualNetworksClient{
		Client: client,
	}, nil
}
