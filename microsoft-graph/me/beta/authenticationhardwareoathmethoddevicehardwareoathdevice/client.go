package authenticationhardwareoathmethoddevicehardwareoathdevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationHardwareOathMethodDeviceHardwareOathDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationHardwareOathMethodDeviceHardwareOathDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationHardwareOathMethodDeviceHardwareOathDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationhardwareoathmethoddevicehardwareoathdevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationHardwareOathMethodDeviceHardwareOathDeviceClient: %+v", err)
	}

	return &AuthenticationHardwareOathMethodDeviceHardwareOathDeviceClient{
		Client: client,
	}, nil
}
