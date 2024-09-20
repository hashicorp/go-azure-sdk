package siteonenotesectiongroupsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionClient{
		Client: client,
	}, nil
}
