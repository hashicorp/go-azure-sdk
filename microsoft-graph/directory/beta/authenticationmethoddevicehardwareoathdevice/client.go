package authenticationmethoddevicehardwareoathdevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodDeviceHardwareOathDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMethodDeviceHardwareOathDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMethodDeviceHardwareOathDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmethoddevicehardwareoathdevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMethodDeviceHardwareOathDeviceClient: %+v", err)
	}

	return &AuthenticationMethodDeviceHardwareOathDeviceClient{
		Client: client,
	}, nil
}
