package authenticationmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmethoddevicehardwareoathdeviceassigntoserviceprovisioningerror", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient: %+v", err)
	}

	return &AuthenticationMethodDeviceHardwareOathDeviceAssignToServiceProvisioningErrorClient{
		Client: client,
	}, nil
}
