package siteonenotesectionparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenoteSectionParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenoteSectionParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenoteSectionParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotesectionparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenoteSectionParentNotebookClient: %+v", err)
	}

	return &SiteOnenoteSectionParentNotebookClient{
		Client: client,
	}, nil
}
