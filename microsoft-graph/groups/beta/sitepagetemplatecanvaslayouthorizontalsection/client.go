package sitepagetemplatecanvaslayouthorizontalsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCanvasLayoutHorizontalSectionClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCanvasLayoutHorizontalSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCanvasLayoutHorizontalSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecanvaslayouthorizontalsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCanvasLayoutHorizontalSectionClient: %+v", err)
	}

	return &SitePageTemplateCanvasLayoutHorizontalSectionClient{
		Client: client,
	}, nil
}
