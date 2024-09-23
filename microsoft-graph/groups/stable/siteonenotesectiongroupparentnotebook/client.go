package siteonenotesectiongroupparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionGroupParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionGroupParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionGroupParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectiongroupparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionGroupParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteSectionGroupParentNotebookClient{
		Client: client,
	}, nil
}
