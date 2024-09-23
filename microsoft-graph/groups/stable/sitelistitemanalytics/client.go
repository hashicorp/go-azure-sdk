package sitelistitemanalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemAnalyticsClient struct {
	Client *msgraph.Client
}

func NewSiteListItemAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemAnalyticsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemanalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemAnalyticsClient: %+v", err)
	}

	return &SiteListItemAnalyticsClient{
		Client: client,
	}, nil
}
