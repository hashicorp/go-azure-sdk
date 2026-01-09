package siteanalyticsitemactivitystatactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsItemActivityStatActivityClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsItemActivityStatActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsItemActivityStatActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticsitemactivitystatactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsItemActivityStatActivityClient: %+v", err)
	}

	return &SiteAnalyticsItemActivityStatActivityClient{
		Client: client,
	}, nil
}
