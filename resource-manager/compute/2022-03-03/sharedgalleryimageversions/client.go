package sharedgalleryimageversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedGalleryImageVersionsClient struct {
	Client *resourcemanager.Client
}

func NewSharedGalleryImageVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*SharedGalleryImageVersionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "sharedgalleryimageversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SharedGalleryImageVersionsClient: %+v", err)
	}

	return &SharedGalleryImageVersionsClient{
		Client: client,
	}, nil
}
