package assignuserauthenticationhardwareoathmethoddevicehardwareoathdeviceto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignUserAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient struct {
	Client *msgraph.Client
}

func NewAssignUserAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClientWithBaseURI(sdkApi sdkEnv.Api) (*AssignUserAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "assignuserauthenticationhardwareoathmethoddevicehardwareoathdeviceto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignUserAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient: %+v", err)
	}

	return &AssignUserAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient{
		Client: client,
	}, nil
}
