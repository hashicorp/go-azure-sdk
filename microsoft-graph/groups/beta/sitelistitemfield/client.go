package sitelistitemfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemFieldClient struct {
	Client *msgraph.Client
}

func NewSiteListItemFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemFieldClient: %+v", err)
	}

	return &SiteListItemFieldClient{
		Client: client,
	}, nil
}
