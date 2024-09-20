package siteonenotesectiongroupsectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupsectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupSectionParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupSectionParentNotebookClient{
		Client: client,
	}, nil
}
