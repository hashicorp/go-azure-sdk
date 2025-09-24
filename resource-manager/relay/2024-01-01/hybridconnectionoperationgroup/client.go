package hybridconnectionoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridConnectionOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewHybridConnectionOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*HybridConnectionOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hybridconnectionoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridConnectionOperationGroupClient: %+v", err)
	}

	return &HybridConnectionOperationGroupClient{
		Client: client,
	}, nil
}
