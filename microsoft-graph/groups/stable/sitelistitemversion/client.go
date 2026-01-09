package sitelistitemversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemVersionClient struct {
	Client *msgraph.Client
}

func NewSiteListItemVersionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemVersionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemVersionClient: %+v", err)
	}

	return &SiteListItemVersionClient{
		Client: client,
	}, nil
}
