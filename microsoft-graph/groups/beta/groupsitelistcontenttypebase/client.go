package groupsitelistcontenttypebase

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListContentTypeBaseClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListContentTypeBaseClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListContentTypeBaseClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcontenttypebase", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListContentTypeBaseClient: %+v", err)
	}

	return &GroupSiteListContentTypeBaseClient{
		Client: client,
	}, nil
}
