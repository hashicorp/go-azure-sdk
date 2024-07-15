package builds

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildsClient struct {
	Client *resourcemanager.Client
}

func NewBuildsClientWithBaseURI(sdkApi sdkEnv.Api) (*BuildsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "builds", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BuildsClient: %+v", err)
	}

	return &BuildsClient{
		Client: client,
	}, nil
}
