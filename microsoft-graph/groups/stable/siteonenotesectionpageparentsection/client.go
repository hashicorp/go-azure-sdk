package siteonenotesectionpageparentsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionPageParentSectionClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionPageParentSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionPageParentSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectionpageparentsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionPageParentSectionClient: %+v", err)
	}

	return &SiteOnenoteSectionPageParentSectionClient{
		Client: client,
	}, nil
}
