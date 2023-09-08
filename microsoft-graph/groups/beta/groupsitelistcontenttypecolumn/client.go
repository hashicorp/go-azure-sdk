package groupsitelistcontenttypecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListContentTypeColumnClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListContentTypeColumnClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListContentTypeColumnClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcontenttypecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListContentTypeColumnClient: %+v", err)
	}

	return &GroupSiteListContentTypeColumnClient{
		Client: client,
	}, nil
}
