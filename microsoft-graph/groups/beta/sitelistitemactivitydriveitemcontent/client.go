package sitelistitemactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewSiteListItemActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemActivityDriveItemContentClient: %+v", err)
	}

	return &SiteListItemActivityDriveItemContentClient{
		Client: client,
	}, nil
}
