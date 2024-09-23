package siteanalyticslastsevenday

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsLastSevenDayClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsLastSevenDayClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsLastSevenDayClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticslastsevenday", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsLastSevenDayClient: %+v", err)
	}

	return &SiteAnalyticsLastSevenDayClient{
		Client: client,
	}, nil
}
