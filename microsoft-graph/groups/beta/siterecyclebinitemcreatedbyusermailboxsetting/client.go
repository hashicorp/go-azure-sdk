package siterecyclebinitemcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinItemCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinItemCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinItemCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinitemcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinItemCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteRecycleBinItemCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
