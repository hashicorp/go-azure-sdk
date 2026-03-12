package processinfooperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProcessInfoOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewProcessInfoOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*ProcessInfoOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "processinfooperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProcessInfoOperationGroupClient: %+v", err)
	}

	return &ProcessInfoOperationGroupClient{
		Client: client,
	}, nil
}
