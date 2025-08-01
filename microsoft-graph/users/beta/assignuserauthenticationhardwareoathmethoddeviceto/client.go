package assignuserauthenticationhardwareoathmethoddeviceto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignUserAuthenticationHardwareOathMethodDeviceToClient struct {
	Client *msgraph.Client
}

func NewAssignUserAuthenticationHardwareOathMethodDeviceToClientWithBaseURI(sdkApi sdkEnv.Api) (*AssignUserAuthenticationHardwareOathMethodDeviceToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "assignuserauthenticationhardwareoathmethoddeviceto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignUserAuthenticationHardwareOathMethodDeviceToClient: %+v", err)
	}

	return &AssignUserAuthenticationHardwareOathMethodDeviceToClient{
		Client: client,
	}, nil
}
