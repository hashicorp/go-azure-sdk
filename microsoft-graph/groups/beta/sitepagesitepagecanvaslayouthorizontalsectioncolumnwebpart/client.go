package sitepagesitepagecanvaslayouthorizontalsectioncolumnwebpart

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageCanvasLayoutHorizontalSectionColumnWebpartClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageCanvasLayoutHorizontalSectionColumnWebpartClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageCanvasLayoutHorizontalSectionColumnWebpartClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagecanvaslayouthorizontalsectioncolumnwebpart", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageCanvasLayoutHorizontalSectionColumnWebpartClient: %+v", err)
	}

	return &SitePageSitePageCanvasLayoutHorizontalSectionColumnWebpartClient{
		Client: client,
	}, nil
}
