package sitelistitemdriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemDriveItemClient struct {
	Client *msgraph.Client
}

func NewSiteListItemDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemdriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemDriveItemClient: %+v", err)
	}

	return &SiteListItemDriveItemClient{
		Client: client,
	}, nil
}
