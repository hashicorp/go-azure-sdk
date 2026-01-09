package sitelistitemdriveitemcontentstream

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemDriveItemContentStreamClient struct {
	Client *msgraph.Client
}

func NewSiteListItemDriveItemContentStreamClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemDriveItemContentStreamClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemdriveitemcontentstream", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemDriveItemContentStreamClient: %+v", err)
	}

	return &SiteListItemDriveItemContentStreamClient{
		Client: client,
	}, nil
}
