package collection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CollectionClient struct {
	Client *resourcemanager.Client
}

func NewCollectionClientWithBaseURI(sdkApi sdkEnv.Api) (*CollectionClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "collection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CollectionClient: %+v", err)
	}

	return &CollectionClient{
		Client: client,
	}, nil
}
