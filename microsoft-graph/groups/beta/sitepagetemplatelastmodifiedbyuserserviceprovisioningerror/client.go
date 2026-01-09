package sitepagetemplatelastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatelastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SitePageTemplateLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
