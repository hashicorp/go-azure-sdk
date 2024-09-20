package siteonenotepagecontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteOnenotePageContentClient struct {
	Client *msgraph.Client
}

func NewSiteOnenotePageContentClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteOnenotePageContentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siteonenotepagecontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteOnenotePageContentClient: %+v", err)
	}

	return &SiteOnenotePageContentClient{
		Client: client,
	}, nil
}
