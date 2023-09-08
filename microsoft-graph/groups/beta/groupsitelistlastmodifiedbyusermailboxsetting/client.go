package groupsitelistlastmodifiedbyusermailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListLastModifiedByUserMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListLastModifiedByUserMailboxSettingClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListLastModifiedByUserMailboxSettingClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistlastmodifiedbyusermailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListLastModifiedByUserMailboxSettingClient: %+v", err)
	}

	return &GroupSiteListLastModifiedByUserMailboxSettingClient{
		Client: client,
	}, nil
}
