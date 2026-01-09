package siteonenotesectiongroupsectionpage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionPageClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionPageClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionPageClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectionpage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionPageClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionPageClient{
		Client: client,
	}, nil
}
