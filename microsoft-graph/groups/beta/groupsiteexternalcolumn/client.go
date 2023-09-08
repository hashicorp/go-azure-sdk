package groupsiteexternalcolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteExternalColumnClient struct {
	Client *msgraph.Client
}

func NewGroupSiteExternalColumnClientWithBaseURI(api sdkEnv.Api) (*GroupSiteExternalColumnClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteexternalcolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteExternalColumnClient: %+v", err)
	}

	return &GroupSiteExternalColumnClient{
		Client: client,
	}, nil
}
