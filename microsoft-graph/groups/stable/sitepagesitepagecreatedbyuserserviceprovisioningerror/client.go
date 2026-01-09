package sitepagesitepagecreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageSitePageCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSitePageSitePageCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageSitePageCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagesitepagecreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageSitePageCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SitePageSitePageCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
