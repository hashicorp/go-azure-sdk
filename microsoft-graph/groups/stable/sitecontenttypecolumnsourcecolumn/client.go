package sitecontenttypecolumnsourcecolumn

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContentTypeColumnSourceColumnClient struct {
	Client *msgraph.Client
}

func NewSiteContentTypeColumnSourceColumnClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteContentTypeColumnSourceColumnClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecontenttypecolumnsourcecolumn", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteContentTypeColumnSourceColumnClient: %+v", err)
	}

	return &SiteContentTypeColumnSourceColumnClient{
		Client: client,
	}, nil
}
