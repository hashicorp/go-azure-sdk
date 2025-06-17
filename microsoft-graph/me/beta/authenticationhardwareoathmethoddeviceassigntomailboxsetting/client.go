package authenticationhardwareoathmethoddeviceassigntomailboxsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationHardwareOathMethodDeviceAssignToMailboxSettingClient struct {
	Client *msgraph.Client
}

func NewAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationHardwareOathMethodDeviceAssignToMailboxSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationhardwareoathmethoddeviceassigntomailboxsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationHardwareOathMethodDeviceAssignToMailboxSettingClient: %+v", err)
	}

	return &AuthenticationHardwareOathMethodDeviceAssignToMailboxSettingClient{
		Client: client,
	}, nil
}
