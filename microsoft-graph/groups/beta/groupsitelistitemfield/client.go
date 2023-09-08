package groupsitelistitemfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemFieldClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemFieldClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemFieldClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemFieldClient: %+v", err)
	}

	return &GroupSiteListItemFieldClient{
		Client: client,
	}, nil
}
