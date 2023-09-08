package groupsitelistcolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListColumnClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListColumnClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListColumnClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListColumnClient: %+v", err)
	}

	return &GroupSiteListColumnClient{
		Client: client,
	}, nil
}
