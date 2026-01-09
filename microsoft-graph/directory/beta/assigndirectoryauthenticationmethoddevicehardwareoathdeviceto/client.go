package assigndirectoryauthenticationmethoddevicehardwareoathdeviceto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient struct {
	Client *msgraph.Client
}

func NewAssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClientWithBaseURI(sdkApi sdkEnv.Api) (*AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "assigndirectoryauthenticationmethoddevicehardwareoathdeviceto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient: %+v", err)
	}

	return &AssignDirectoryAuthenticationMethodDeviceHardwareOathDeviceToClient{
		Client: client,
	}, nil
}
