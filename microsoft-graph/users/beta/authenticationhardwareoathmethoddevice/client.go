package authenticationhardwareoathmethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationHardwareOathMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationHardwareOathMethodDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationHardwareOathMethodDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationhardwareoathmethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationHardwareOathMethodDeviceClient: %+v", err)
	}

	return &AuthenticationHardwareOathMethodDeviceClient{
		Client: client,
	}, nil
}
