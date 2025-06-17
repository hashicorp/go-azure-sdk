package sitepagetemplatelastmodifiedbyuser

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateLastModifiedByUserClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateLastModifiedByUserClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateLastModifiedByUserClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatelastmodifiedbyuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateLastModifiedByUserClient: %+v", err)
	}

	return &SitePageTemplateLastModifiedByUserClient{
		Client: client,
	}, nil
}
