package sitepagetemplatecreatedbyuserserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SitePageTemplateCreatedByUserServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewSitePageTemplateCreatedByUserServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*SitePageTemplateCreatedByUserServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "sitepagetemplatecreatedbyuserserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SitePageTemplateCreatedByUserServiceProvisioningErrorClient: %+v", err)
	}

	return &SitePageTemplateCreatedByUserServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
