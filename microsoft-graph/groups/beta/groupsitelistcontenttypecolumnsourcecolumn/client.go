package groupsitelistcontenttypecolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListContentTypeColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListContentTypeColumnSourceColumnClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListContentTypeColumnSourceColumnClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcontenttypecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListContentTypeColumnSourceColumnClient: %+v", err)
	}

	return &GroupSiteListContentTypeColumnSourceColumnClient{
		Client: client,
	}, nil
}
