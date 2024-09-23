package sitelistcreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteListCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteListCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteListCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelistcreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteListCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteListCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
