package builders

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuildersClient struct {
	Client *resourcemanager.Client
}

func NewBuildersClientWithBaseURI(sdkApi sdkEnv.Api) (*BuildersClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "builders", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BuildersClient: %+v", err)
	}

	return &BuildersClient{
		Client: client,
	}, nil
}
