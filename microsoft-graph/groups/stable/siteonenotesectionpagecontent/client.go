package siteonenotesectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionPageContentClient: %+v", err)
	}

	return &SiteOnenoteSectionPageContentClient{
		Client: client,
	}, nil
}
