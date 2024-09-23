package sitelistitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemClient struct {
	Client *msgraph.Client
}

func NewSiteListItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemClient: %+v", err)
	}

	return &SiteListItemClient{
		Client: client,
	}, nil
}
