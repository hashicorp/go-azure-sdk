package contentpackages

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentPackagesClient struct {
	Client *resourcemanager.Client
}

func NewContentPackagesClientWithBaseURI(sdkApi sdkEnv.Api) (*ContentPackagesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "contentpackages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ContentPackagesClient: %+v", err)
	}

	return &ContentPackagesClient{
		Client: client,
	}, nil
}
