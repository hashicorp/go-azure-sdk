package sitecontenttypecolumnposition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContentTypeColumnPositionClient struct {
	Client *msgraph.Client
}

func NewSiteContentTypeColumnPositionClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteContentTypeColumnPositionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecontenttypecolumnposition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteContentTypeColumnPositionClient: %+v", err)
	}

	return &SiteContentTypeColumnPositionClient{
		Client: client,
	}, nil
}
