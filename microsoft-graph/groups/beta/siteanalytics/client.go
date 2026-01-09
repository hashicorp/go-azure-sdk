package siteanalytics

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalytics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsClient: %+v", err)
	}

	return &SiteAnalyticsClient{
		Client: client,
	}, nil
}
