package sitecontenttype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContentTypeClient struct {
	Client *msgraph.Client
}

func NewSiteContentTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteContentTypeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecontenttype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteContentTypeClient: %+v", err)
	}

	return &SiteContentTypeClient{
		Client: client,
	}, nil
}
