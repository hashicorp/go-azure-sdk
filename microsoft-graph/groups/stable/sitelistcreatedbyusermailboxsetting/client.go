package sitelistcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteListCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteListCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
