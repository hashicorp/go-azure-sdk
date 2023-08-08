package sharedgalleries

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleriesClient struct {
	Client *resourcemanager.Client
}

func NewSharedGalleriesClientWithBaseURI(sdkApi sdkEnv.Api) (*SharedGalleriesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sharedgalleries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SharedGalleriesClient: %+v", err)
	}

	return &SharedGalleriesClient{
		Client: client,
	}, nil
}
