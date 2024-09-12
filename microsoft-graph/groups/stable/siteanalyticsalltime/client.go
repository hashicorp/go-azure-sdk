package siteanalyticsalltime

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsAllTimeClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsAllTimeClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsAllTimeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticsalltime", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsAllTimeClient: %+v", err)
	}

	return &SiteAnalyticsAllTimeClient{
		Client: client,
	}, nil
}
