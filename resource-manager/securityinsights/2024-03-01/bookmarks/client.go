package bookmarks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarksClient struct {
	Client *resourcemanager.Client
}

func NewBookmarksClientWithBaseURI(sdkApi sdkEnv.Api) (*BookmarksClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "bookmarks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BookmarksClient: %+v", err)
	}

	return &BookmarksClient{
		Client: client,
	}, nil
}
