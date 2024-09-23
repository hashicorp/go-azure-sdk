package sitecreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewSiteCreatedByUserMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &SiteCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
