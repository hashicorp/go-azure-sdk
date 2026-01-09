package authenticationhardwareoathmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewAuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationhardwareoathmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient: %+v", err)
	}

	return &AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
