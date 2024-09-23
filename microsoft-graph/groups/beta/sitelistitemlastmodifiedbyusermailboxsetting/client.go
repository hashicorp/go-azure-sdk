package sitelistitemlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteListItemLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteListItemLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
