package siteanalyticitemactivitystatactivity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticItemActivityStatActivityClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticItemActivityStatActivityClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticItemActivityStatActivityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticitemactivitystatactivity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticItemActivityStatActivityClient: %+v", err)
	}

	return &SiteAnalyticItemActivityStatActivityClient{
		Client: client,
	}, nil
}
