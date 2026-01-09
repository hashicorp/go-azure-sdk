package siterecyclebincreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebincreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteRecycleBinCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
