package sitelistcontenttypebase

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeBaseClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeBaseClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeBaseClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttypebase", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeBaseClient: %+v", err)
	}

	return &SiteListContentTypeBaseClient{
		Client: client,
	}, nil
}
