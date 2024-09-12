package sitelastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSiteLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SiteLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitelastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SiteLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SiteLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
