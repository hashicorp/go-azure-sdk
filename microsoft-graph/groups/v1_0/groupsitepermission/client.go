package groupsitepermission

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSitePermissionClient struct {
	Client *msgraph.Client
}

func NewGroupSitePermissionClientWithBaseURI(api sdkEnv.Api) (*GroupSitePermissionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitepermission", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSitePermissionClient: %+v", err)
	}

	return &GroupSitePermissionClient{
		Client: client,
	}, nil
}
