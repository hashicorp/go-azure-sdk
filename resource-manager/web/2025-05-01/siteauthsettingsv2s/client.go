package siteauthsettingsv2s

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAuthSettingsV2sClient struct {
	Client *resourcemanager.Client
}

func NewSiteAuthSettingsV2sClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAuthSettingsV2sClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "siteauthsettingsv2s", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAuthSettingsV2sClient: %+v", err)
	}

	return &SiteAuthSettingsV2sClient{
		Client: client,
	}, nil
}
