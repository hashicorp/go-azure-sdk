package groupsitelistcreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListCreatedByUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &GroupSiteListCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
