package communitygalleryimageversions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImageVersionsClient struct {
	Client *resourcemanager.Client
}

func NewCommunityGalleryImageVersionsClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunityGalleryImageVersionsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "communitygalleryimageversions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunityGalleryImageVersionsClient: %+v", err)
	}

	return &CommunityGalleryImageVersionsClient{
		Client: client,
	}, nil
}
