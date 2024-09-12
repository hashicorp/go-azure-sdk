package mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewMobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobiledevicemanagementpolicyincludedgroupserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient: %+v", err)
	}

	return &MobileDeviceManagementPolicyIncludedGroupServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
