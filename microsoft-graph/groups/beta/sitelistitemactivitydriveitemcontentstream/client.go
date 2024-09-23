package sitelistitemactivitydriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemActivityDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewSiteListItemActivityDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemActivityDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemactivitydriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemActivityDriveItemContentStreamClient: %+v", err)
	}

	return &SiteListItemActivityDriveItemContentStreamClient{
		Client: client,
	}, nil
}
