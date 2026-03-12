package swiftvirtualnetworkoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SwiftVirtualNetworkOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewSwiftVirtualNetworkOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SwiftVirtualNetworkOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "swiftvirtualnetworkoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SwiftVirtualNetworkOperationGroupClient: %+v", err)
	}

	return &SwiftVirtualNetworkOperationGroupClient{
		Client: client,
	}, nil
}
