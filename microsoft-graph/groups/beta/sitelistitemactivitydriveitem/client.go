package sitelistitemactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewSiteListItemActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemActivityDriveItemClient: %+v", err)
	}

	return &SiteListItemActivityDriveItemClient{
		Client: client,
	}, nil
}
