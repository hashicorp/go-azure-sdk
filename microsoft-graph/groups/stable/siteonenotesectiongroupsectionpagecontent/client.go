package siteonenotesectiongroupsectionpagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionPageContentClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionPageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionPageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectionpagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionPageContentClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionPageContentClient{
		Client: client,
	}, nil
}
