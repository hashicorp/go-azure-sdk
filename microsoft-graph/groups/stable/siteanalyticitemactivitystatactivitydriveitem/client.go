package siteanalyticitemactivitystatactivitydriveitem

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticItemActivityStatActivityDriveItemClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticItemActivityStatActivityDriveItemClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticItemActivityStatActivityDriveItemClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticitemactivitystatactivitydriveitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticItemActivityStatActivityDriveItemClient: %+v", err)
	}

	return &SiteAnalyticItemActivityStatActivityDriveItemClient{
		Client: client,
	}, nil
}
