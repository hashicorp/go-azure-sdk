package authenticationhardwareoathmethoddevicehardwareoathdeviceassigntomailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewAuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationhardwareoathmethoddevicehardwareoathdeviceassigntomailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient: %+v", err)
	}

	return &AuthenticationHardwareOathMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient{
		Client: client,
	}, nil
}
