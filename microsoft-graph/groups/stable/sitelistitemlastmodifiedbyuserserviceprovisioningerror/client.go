package sitelistitemlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListItemLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteListItemLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListItemLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistitemlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListItemLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteListItemLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
