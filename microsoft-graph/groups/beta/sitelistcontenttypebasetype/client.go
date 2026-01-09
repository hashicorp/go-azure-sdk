package sitelistcontenttypebasetype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeBaseTypeClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeBaseTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeBaseTypeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttypebasetype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeBaseTypeClient: %+v", err)
	}

	return &SiteListContentTypeBaseTypeClient{
		Client: client,
	}, nil
}
