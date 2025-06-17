package sitepagetemplatecanvaslayout

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCanvasLayoutClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCanvasLayoutClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCanvasLayoutClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecanvaslayout", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCanvasLayoutClient: %+v", err)
	}

	return &SitePageTemplateCanvasLayoutClient{
		Client: client,
	}, nil
}
