package groupsitelistitemactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemActivityDriveItemClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemActivityDriveItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemActivityDriveItemClient: %+v", err)
	}

	return &GroupSiteListItemActivityDriveItemClient{
		Client: client,
	}, nil
}
