package sitelistlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteListLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteListLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
