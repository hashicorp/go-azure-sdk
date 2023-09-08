package groupsitelistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemClient: %+v", err)
	}

	return &GroupSiteListItemClient{
		Client: client,
	}, nil
}
