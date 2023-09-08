package groupsitelistitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemActivityClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemActivityClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemActivityClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemActivityClient: %+v", err)
	}

	return &GroupSiteListItemActivityClient{
		Client: client,
	}, nil
}
