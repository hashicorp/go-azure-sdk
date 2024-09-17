package d4apicollectionlist

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type D4APICollectionListClient struct {
	Client *resourcemanager.Client
}

func NewD4APICollectionListClientWithBaseURI(sdkApi sdkEnv.Api) (*D4APICollectionListClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "d4apicollectionlist", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating D4APICollectionListClient: %+v", err)
	}

	return &D4APICollectionListClient{
		Client: client,
	}, nil
}
