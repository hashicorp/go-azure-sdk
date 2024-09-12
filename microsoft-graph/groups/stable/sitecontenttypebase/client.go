package sitecontenttypebase

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContentTypeBaseClient struct {
	Client *msgraph.Client
}

func NewSiteContentTypeBaseClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteContentTypeBaseClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecontenttypebase", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteContentTypeBaseClient: %+v", err)
	}

	return &SiteContentTypeBaseClient{
		Client: client,
	}, nil
}
