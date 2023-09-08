package groupsitecreatedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteCreatedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewGroupSiteCreatedByUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*GroupSiteCreatedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitecreatedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteCreatedByUserMailboxSettingClient: %+v", err)
	}

	return &GroupSiteCreatedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
