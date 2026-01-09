package sitelistcontenttypecolumnlink

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeColumnLinkClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeColumnLinkClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeColumnLinkClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttypecolumnlink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeColumnLinkClient: %+v", err)
	}

	return &SiteListContentTypeColumnLinkClient{
		Client: client,
	}, nil
}
