package siteanalyticsitemactivitystatactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsItemActivityStatActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsItemActivityStatActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticsitemactivitystatactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsItemActivityStatActivityDriveItemClient: %+v", err)
	}

	return &SiteAnalyticsItemActivityStatActivityDriveItemClient{
		Client: client,
	}, nil
}
