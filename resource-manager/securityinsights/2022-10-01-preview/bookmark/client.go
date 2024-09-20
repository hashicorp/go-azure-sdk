package bookmark

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookmarkClient struct {
	Client *resourcemanager.Client
}

func NewBookmarkClientWithBaseURI(sdkApi sdkEnv.Api) (*BookmarkClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "bookmark", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BookmarkClient: %+v", err)
	}

	return &BookmarkClient{
		Client: client,
	}, nil
}
