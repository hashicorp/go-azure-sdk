package sitelistitemdriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewSiteListItemDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemdriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemDriveItemContentClient: %+v", err)
	}

	return &SiteListItemDriveItemContentClient{
		Client: client,
	}, nil
}
