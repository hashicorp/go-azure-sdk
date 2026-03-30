package bmckeysets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BmcKeySetsClient struct {
	Client *resourcemanager.Client
}

func NewBmcKeySetsClientWithBaseURI(sdkApi sdkEnv.Api) (*BmcKeySetsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "bmckeysets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BmcKeySetsClient: %+v", err)
	}

	return &BmcKeySetsClient{
		Client: client,
	}, nil
}
