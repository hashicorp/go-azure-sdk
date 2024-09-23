package sitelistcontenttype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListContentTypeClient struct {
	Client *msgraph.Client
}

func NewSiteListContentTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListContentTypeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcontenttype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListContentTypeClient: %+v", err)
	}

	return &SiteListContentTypeClient{
		Client: client,
	}, nil
}
