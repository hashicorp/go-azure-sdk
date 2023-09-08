package groupsitelistcontenttypecolumnlink

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListContentTypeColumnLinkClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListContentTypeColumnLinkClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListContentTypeColumnLinkClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcontenttypecolumnlink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListContentTypeColumnLinkClient: %+v", err)
	}

	return &GroupSiteListContentTypeColumnLinkClient{
		Client: client,
	}, nil
}
