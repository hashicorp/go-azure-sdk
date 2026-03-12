package siteextensioninfos

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteExtensionInfosClient struct {
	Client *resourcemanager.Client
}

func NewSiteExtensionInfosClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteExtensionInfosClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "siteextensioninfos", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteExtensionInfosClient: %+v", err)
	}

	return &SiteExtensionInfosClient{
		Client: client,
	}, nil
}
