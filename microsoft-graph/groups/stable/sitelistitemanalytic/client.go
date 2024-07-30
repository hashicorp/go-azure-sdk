package sitelistitemanalytic

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemAnalyticClient struct {
	Client *msgraph.Client
}

func NewSiteListItemAnalyticClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemAnalyticClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemanalytic", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemAnalyticClient: %+v", err)
	}

	return &SiteListItemAnalyticClient{
		Client: client,
	}, nil
}
