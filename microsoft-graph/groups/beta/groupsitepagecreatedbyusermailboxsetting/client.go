package groupsitepagecreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSitePageCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewGroupSitePageCreatedByUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*GroupSitePageCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitepagecreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSitePageCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &GroupSitePageCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
