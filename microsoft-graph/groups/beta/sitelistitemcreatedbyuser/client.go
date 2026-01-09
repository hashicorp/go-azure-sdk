package sitelistitemcreatedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemCreatedByUserClient struct {
	Client *msgraph.Client
}

func NewSiteListItemCreatedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemCreatedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemcreatedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemCreatedByUserClient: %+v", err)
	}

	return &SiteListItemCreatedByUserClient{
		Client: client,
	}, nil
}
