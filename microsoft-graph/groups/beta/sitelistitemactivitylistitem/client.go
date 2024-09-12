package sitelistitemactivitylistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemActivityListItemClient struct {
	Client *msgraph.Client
}

func NewSiteListItemActivityListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemActivityListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemactivitylistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemActivityListItemClient: %+v", err)
	}

	return &SiteListItemActivityListItemClient{
		Client: client,
	}, nil
}
