package sitelistcontenttypecolumnposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeColumnPositionClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeColumnPositionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeColumnPositionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttypecolumnposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeColumnPositionClient: %+v", err)
	}

	return &SiteListContentTypeColumnPositionClient{
		Client: client,
	}, nil
}
