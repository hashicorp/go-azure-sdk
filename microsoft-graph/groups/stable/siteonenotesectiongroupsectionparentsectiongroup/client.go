package siteonenotesectiongroupsectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionParentSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
