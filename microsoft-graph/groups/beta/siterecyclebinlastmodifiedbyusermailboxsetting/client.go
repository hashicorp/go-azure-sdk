package siterecyclebinlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteRecycleBinLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
