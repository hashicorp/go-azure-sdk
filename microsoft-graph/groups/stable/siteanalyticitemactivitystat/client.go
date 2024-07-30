package siteanalyticitemactivitystat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteAnalyticItemActivityStatClient struct {
	Client *msgraph.Client
}

func NewSiteAnalyticItemActivityStatClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteAnalyticItemActivityStatClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteanalyticitemactivitystat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteAnalyticItemActivityStatClient: %+v", err)
	}

	return &SiteAnalyticItemActivityStatClient{
		Client: client,
	}, nil
}
