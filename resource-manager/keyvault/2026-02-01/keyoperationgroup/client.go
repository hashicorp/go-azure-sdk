package keyoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewKeyOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*KeyOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "keyoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KeyOperationGroupClient: %+v", err)
	}

	return &KeyOperationGroupClient{
		Client: client,
	}, nil
}
