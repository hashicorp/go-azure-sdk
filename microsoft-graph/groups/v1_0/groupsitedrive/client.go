package groupsitedrive

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteDriveClient struct {
	Client *msgraph.Client
}

func NewGroupSiteDriveClientWithBaseURI(api sdkEnv.Api) (*GroupSiteDriveClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitedrive", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteDriveClient: %+v", err)
	}

	return &GroupSiteDriveClient{
		Client: client,
	}, nil
}
