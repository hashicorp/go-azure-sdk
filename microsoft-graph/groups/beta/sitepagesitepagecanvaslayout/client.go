package sitepagesitepagecanvaslayout

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageCanvasLayoutClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageCanvasLayoutClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageCanvasLayoutClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagecanvaslayout", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageCanvasLayoutClient: %+v", err)
	}

	return &SitePageSitePageCanvasLayoutClient{
		Client: client,
	}, nil
}
