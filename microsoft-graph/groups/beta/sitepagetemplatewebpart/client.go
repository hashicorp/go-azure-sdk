package sitepagetemplatewebpart

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateWebPartClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateWebPartClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateWebPartClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatewebpart", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateWebPartClient: %+v", err)
	}

	return &SitePageTemplateWebPartClient{
		Client: client,
	}, nil
}
