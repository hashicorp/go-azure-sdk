package hybridconnectionlimitsoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HybridConnectionLimitsOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewHybridConnectionLimitsOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*HybridConnectionLimitsOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "hybridconnectionlimitsoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HybridConnectionLimitsOperationGroupClient: %+v", err)
	}

	return &HybridConnectionLimitsOperationGroupClient{
		Client: client,
	}, nil
}
