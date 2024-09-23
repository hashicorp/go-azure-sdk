package siteanalyticsitemactivitystat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsItemActivityStatClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsItemActivityStatClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsItemActivityStatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticsitemactivitystat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsItemActivityStatClient: %+v", err)
	}

	return &SiteAnalyticsItemActivityStatClient{
		Client: client,
	}, nil
}
