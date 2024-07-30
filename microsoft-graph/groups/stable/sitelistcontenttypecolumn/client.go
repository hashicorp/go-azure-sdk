package sitelistcontenttypecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeColumnClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttypecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeColumnClient: %+v", err)
	}

	return &SiteListContentTypeColumnClient{
		Client: client,
	}, nil
}
