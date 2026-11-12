package dppbaseresourceoperationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DppBaseResourceOperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewDppBaseResourceOperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*DppBaseResourceOperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "dppbaseresourceoperationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DppBaseResourceOperationGroupClient: %+v", err)
	}

	return &DppBaseResourceOperationGroupClient{
		Client: client,
	}, nil
}
