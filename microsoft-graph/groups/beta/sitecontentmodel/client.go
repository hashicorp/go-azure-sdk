package sitecontentmodel

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContentModelClient struct {
	Client *msgraph.Client
}

func NewSiteContentModelClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteContentModelClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecontentmodel", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteContentModelClient: %+v", err)
	}

	return &SiteContentModelClient{
		Client: client,
	}, nil
}
