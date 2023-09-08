package groupsitelistitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemCreatedByUserClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemCreatedByUserClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemCreatedByUserClient: %+v", err)
	}

	return &GroupSiteListItemCreatedByUserClient{
		Client: client,
	}, nil
}
