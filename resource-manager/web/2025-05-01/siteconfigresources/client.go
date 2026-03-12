package siteconfigresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteConfigResourcesClient struct {
	Client *resourcemanager.Client
}

func NewSiteConfigResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteConfigResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "siteconfigresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteConfigResourcesClient: %+v", err)
	}

	return &SiteConfigResourcesClient{
		Client: client,
	}, nil
}
