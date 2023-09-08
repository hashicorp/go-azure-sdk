package groupsitelistitemlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemLastModifiedByUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &GroupSiteListItemLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
