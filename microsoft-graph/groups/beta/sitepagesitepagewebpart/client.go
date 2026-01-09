package sitepagesitepagewebpart

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageWebPartClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageWebPartClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageWebPartClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagewebpart", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageWebPartClient: %+v", err)
	}

	return &SitePageSitePageWebPartClient{
		Client: client,
	}, nil
}
