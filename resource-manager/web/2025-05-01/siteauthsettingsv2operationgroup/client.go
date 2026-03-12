package siteauthsettingsv2operationgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAuthSettingsV2OperationGroupClient struct {
	Client *resourcemanager.Client
}

func NewSiteAuthSettingsV2OperationGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAuthSettingsV2OperationGroupClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "siteauthsettingsv2operationgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAuthSettingsV2OperationGroupClient: %+v", err)
	}

	return &SiteAuthSettingsV2OperationGroupClient{
		Client: client,
	}, nil
}
