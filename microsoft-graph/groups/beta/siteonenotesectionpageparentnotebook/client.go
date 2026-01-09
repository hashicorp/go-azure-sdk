package siteonenotesectionpageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionPageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionPageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionPageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectionpageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionPageParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteSectionPageParentNotebookClient{
		Client: client,
	}, nil
}
