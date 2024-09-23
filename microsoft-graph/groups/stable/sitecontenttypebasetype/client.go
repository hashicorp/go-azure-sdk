package sitecontenttypebasetype

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteContentTypeBaseTypeClient struct {
	Client *msgraph.Client
}

func NewSiteContentTypeBaseTypeClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteContentTypeBaseTypeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitecontenttypebasetype", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteContentTypeBaseTypeClient: %+v", err)
	}

	return &SiteContentTypeBaseTypeClient{
		Client: client,
	}, nil
}
