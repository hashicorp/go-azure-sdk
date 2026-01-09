package siterecyclebincreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebincreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteRecycleBinCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
