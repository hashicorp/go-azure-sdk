package sitelogsconfigs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteLogsConfigsClient struct {
	Client *resourcemanager.Client
}

func NewSiteLogsConfigsClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteLogsConfigsClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "sitelogsconfigs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteLogsConfigsClient: %+v", err)
	}

	return &SiteLogsConfigsClient{
		Client: client,
	}, nil
}
