package groupsitelistitemdriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemDriveItemClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemDriveItemClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemDriveItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemdriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemDriveItemClient: %+v", err)
	}

	return &GroupSiteListItemDriveItemClient{
		Client: client,
	}, nil
}
