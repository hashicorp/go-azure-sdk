package mobileappmanagementpolicyincludedgroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewMobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobileappmanagementpolicyincludedgroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &MobileAppManagementPolicyIncludedGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
