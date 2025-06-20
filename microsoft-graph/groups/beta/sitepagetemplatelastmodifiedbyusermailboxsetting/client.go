package sitepagetemplatelastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateLastModifiedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatelastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &SitePageTemplateLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
