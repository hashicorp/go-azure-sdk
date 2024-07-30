package siteanalyticlastsevenday

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticLastSevenDayClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticLastSevenDayClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticLastSevenDayClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticlastsevenday", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticLastSevenDayClient: %+v", err)
	}

	return &SiteAnalyticLastSevenDayClient{
		Client: client,
	}, nil
}
