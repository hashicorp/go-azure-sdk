package siteonenotesectionparentsectiongroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionParentSectionGroupClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionParentSectionGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionParentSectionGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectionparentsectiongroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionParentSectionGroupClient: %+v", err)
	}

	return &SiteOnenoteSectionParentSectionGroupClient{
		Client: client,
	}, nil
}
