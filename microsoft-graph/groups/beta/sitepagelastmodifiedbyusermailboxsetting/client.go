package sitepagelastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSitePageLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagelastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SitePageLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
