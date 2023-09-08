package groupsitelistitemactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListItemActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListItemActivityDriveItemContentClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListItemActivityDriveItemContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistitemactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListItemActivityDriveItemContentClient: %+v", err)
	}

	return &GroupSiteListItemActivityDriveItemContentClient{
		Client: client,
	}, nil
}
