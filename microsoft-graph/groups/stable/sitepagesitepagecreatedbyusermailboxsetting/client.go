package sitepagesitepagecreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagecreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SitePageSitePageCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
