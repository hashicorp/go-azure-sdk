package groupsitelistdrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListDriveClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListDriveClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListDriveClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistdrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListDriveClient: %+v", err)
	}

	return &GroupSiteListDriveClient{
		Client: client,
	}, nil
}
