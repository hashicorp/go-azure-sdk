package groupsiteitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteItemClient struct {
	Client *msgraph.Client
}

func NewGroupSiteItemClientWithBaseURI(api sdkEnv.Api) (*GroupSiteItemClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsiteitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteItemClient: %+v", err)
	}

	return &GroupSiteItemClient{
		Client: client,
	}, nil
}
