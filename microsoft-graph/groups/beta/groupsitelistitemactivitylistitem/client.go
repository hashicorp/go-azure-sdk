package groupsitelistitemactivitylistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemActivityListItemClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemActivityListItemClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemActivityListItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemActivityListItemClient: %+v", err)
	}

	return &GroupSiteListItemActivityListItemClient{
		Client: client,
	}, nil
}
