package siterecyclebinitemcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteRecycleBinItemCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteRecycleBinItemCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteRecycleBinItemCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "siterecyclebinitemcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteRecycleBinItemCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteRecycleBinItemCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
