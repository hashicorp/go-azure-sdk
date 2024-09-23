package sitepagecreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSitePageCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagecreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SitePageCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
