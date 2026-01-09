package siterecyclebinlastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinlastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteRecycleBinLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
