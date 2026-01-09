package sitepagetemplatecanvaslayouthorizontalsectioncolumnwebpart

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecanvaslayouthorizontalsectioncolumnwebpart", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartClient: %+v", err)
	}

	return &SitePageTemplateCanvasLayoutHorizontalSectionColumnWebpartClient{
		Client: client,
	}, nil
}
