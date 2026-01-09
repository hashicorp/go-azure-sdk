package sitepagesitepagelastmodifiedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageLastModifiedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageLastModifiedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageLastModifiedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagelastmodifiedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageLastModifiedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SitePageSitePageLastModifiedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
