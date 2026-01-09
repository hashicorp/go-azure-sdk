package sitelistitemactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemActivityClient struct {
	Client *msgraph.Client
}

func NewSiteListItemActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemActivityClient: %+v", err)
	}

	return &SiteListItemActivityClient{
		Client: client,
	}, nil
}
