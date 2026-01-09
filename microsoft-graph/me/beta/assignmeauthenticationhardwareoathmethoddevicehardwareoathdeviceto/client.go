package assignmeauthenticationhardwareoathmethoddevicehardwareoathdeviceto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignMeAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient struct {
	Client *msgraph.Client
}

func NewAssignMeAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClientWithBaseURI(sdkApi sdkEnv.Api) (*AssignMeAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "assignmeauthenticationhardwareoathmethoddevicehardwareoathdeviceto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignMeAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient: %+v", err)
	}

	return &AssignMeAuthenticationHardwareOathMethodDeviceHardwareOathDeviceToClient{
		Client: client,
	}, nil
}
