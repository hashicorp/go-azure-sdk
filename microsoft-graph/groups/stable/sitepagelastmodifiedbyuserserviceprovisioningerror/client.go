package sitepagelastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSitePageLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagelastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SitePageLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
