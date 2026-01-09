package siteonenotesectiongroupsectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionPageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionPageParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
