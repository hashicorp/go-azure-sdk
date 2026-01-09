package sitepagetemplate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateClient: %+v", err)
	}

	return &SitePageTemplateClient{
		Client: client,
	}, nil
}
