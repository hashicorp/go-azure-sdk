package bookmarkrelations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarkRelationsClient struct {
	Client *resourcemanager.Client
}

func NewBookmarkRelationsClientWithBaseURI(sdkApi sdkEnv.Api) (*BookmarkRelationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "bookmarkrelations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BookmarkRelationsClient: %+v", err)
	}

	return &BookmarkRelationsClient{
		Client: client,
	}, nil
}
