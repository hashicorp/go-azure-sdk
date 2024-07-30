package siteanalyticitemactivitystatactivitydriveitemcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticItemActivityStatActivityDriveItemContentClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticItemActivityStatActivityDriveItemContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticItemActivityStatActivityDriveItemContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticitemactivitystatactivitydriveitemcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticItemActivityStatActivityDriveItemContentClient: %+v", err)
	}

	return &SiteAnalyticItemActivityStatActivityDriveItemContentClient{
		Client: client,
	}, nil
}
