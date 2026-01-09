package sitelistitemversionfield

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemVersionFieldClient struct {
	Client *msgraph.Client
}

func NewSiteListItemVersionFieldClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemVersionFieldClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemversionfield", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemVersionFieldClient: %+v", err)
	}

	return &SiteListItemVersionFieldClient{
		Client: client,
	}, nil
}
