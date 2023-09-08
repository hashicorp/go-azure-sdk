package groupsitelistcontenttypebasetype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupSiteListContentTypeBaseTypeClient struct {
	Client *msgraph.Client
}

func NewGroupSiteListContentTypeBaseTypeClientWithBaseURI(api sdkEnv.Api) (*GroupSiteListContentTypeBaseTypeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupsitelistcontenttypebasetype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupSiteListContentTypeBaseTypeClient: %+v", err)
	}

	return &GroupSiteListContentTypeBaseTypeClient{
		Client: client,
	}, nil
}
