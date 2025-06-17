package sitepagetemplatecanvaslayouthorizontalsectioncolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCanvasLayoutHorizontalSectionColumnClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCanvasLayoutHorizontalSectionColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCanvasLayoutHorizontalSectionColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecanvaslayouthorizontalsectioncolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCanvasLayoutHorizontalSectionColumnClient: %+v", err)
	}

	return &SitePageTemplateCanvasLayoutHorizontalSectionColumnClient{
		Client: client,
	}, nil
}
