package groupsitelistitemversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemVersionClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemVersionClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemVersionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemVersionClient: %+v", err)
	}

	return &GroupSiteListItemVersionClient{
		Client: client,
	}, nil
}
