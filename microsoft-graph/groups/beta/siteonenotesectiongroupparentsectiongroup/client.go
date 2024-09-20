package siteonenotesectiongroupparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupParentSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupParentSectionGroupClient{
		Client: client,
	}, nil
}
