package d4apicollection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type D4APICollectionClient struct {
	Client *resourcemanager.Client
}

func NewD4APICollectionClientWithBaseURI(sdkApi sdkEnv.Api) (*D4APICollectionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "d4apicollection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating D4APICollectionClient: %+v", err)
	}

	return &D4APICollectionClient{
		Client: client,
	}, nil
}
