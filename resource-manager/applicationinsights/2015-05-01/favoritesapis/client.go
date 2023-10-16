package favoritesapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FavoritesAPIsClient struct {
	Client *resourcemanager.Client
}

func NewFavoritesAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*FavoritesAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "favoritesapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FavoritesAPIsClient: %+v", err)
	}

	return &FavoritesAPIsClient{
		Client: client,
	}, nil
}
