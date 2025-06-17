package sitepagetemplatecanvaslayoutverticalsection

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCanvasLayoutVerticalSectionClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCanvasLayoutVerticalSectionClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCanvasLayoutVerticalSectionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecanvaslayoutverticalsection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCanvasLayoutVerticalSectionClient: %+v", err)
	}

	return &SitePageTemplateCanvasLayoutVerticalSectionClient{
		Client: client,
	}, nil
}
