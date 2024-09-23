package siteonenotepageparentnotebook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenotePageParentNotebookClient struct {
	Client *msgraph.Client
}

func NewSiteOnenotePageParentNotebookClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenotePageParentNotebookClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotepageparentnotebook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenotePageParentNotebookClient: %+v", err)
	}

	return &SiteOnenotePageParentNotebookClient{
		Client: client,
	}, nil
}
