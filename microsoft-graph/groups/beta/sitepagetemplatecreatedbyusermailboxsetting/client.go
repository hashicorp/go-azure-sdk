package sitepagetemplatecreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SitePageTemplateCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
