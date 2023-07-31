package videos

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VideosClient struct {
	Client *resourcemanager.Client
}

func NewVideosClientWithBaseURI(api sdkEnv.Api) (*VideosClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "videos", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VideosClient: %+v", err)
	}

	return &VideosClient{
		Client: client,
	}, nil
}
