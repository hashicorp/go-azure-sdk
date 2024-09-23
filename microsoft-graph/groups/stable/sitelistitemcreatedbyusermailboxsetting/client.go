package sitelistitemcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteListItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteListItemCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
