package sitepagesitepagecanvaslayouthorizontalsectioncolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageCanvasLayoutHorizontalSectionColumnClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageCanvasLayoutHorizontalSectionColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageCanvasLayoutHorizontalSectionColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagecanvaslayouthorizontalsectioncolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageCanvasLayoutHorizontalSectionColumnClient: %+v", err)
	}

	return &SitePageSitePageCanvasLayoutHorizontalSectionColumnClient{
		Client: client,
	}, nil
}
