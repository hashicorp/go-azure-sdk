package siteonenotesectiongroupsectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionPageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionPageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionPageParentSectionClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionPageParentSectionClient{
		Client: client,
	}, nil
}
