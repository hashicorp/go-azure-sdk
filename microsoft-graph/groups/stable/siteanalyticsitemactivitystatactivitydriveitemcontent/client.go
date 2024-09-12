package siteanalyticsitemactivitystatactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticsItemActivityStatActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticsItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticsItemActivityStatActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticsitemactivitystatactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticsItemActivityStatActivityDriveItemContentClient: %+v", err)
	}

	return &SiteAnalyticsItemActivityStatActivityDriveItemContentClient{
		Client: client,
	}, nil
}
