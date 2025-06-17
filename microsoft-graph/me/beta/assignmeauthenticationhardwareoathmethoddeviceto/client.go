package assignmeauthenticationhardwareoathmethoddeviceto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignMeAuthenticationHardwareOathMethodDeviceToClient struct {
	Client *msgraph.Client
}

func NewAssignMeAuthenticationHardwareOathMethodDeviceToClientWithBaseURI(sdkApi sdkEnv.Api) (*AssignMeAuthenticationHardwareOathMethodDeviceToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "assignmeauthenticationhardwareoathmethoddeviceto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AssignMeAuthenticationHardwareOathMethodDeviceToClient: %+v", err)
	}

	return &AssignMeAuthenticationHardwareOathMethodDeviceToClient{
		Client: client,
	}, nil
}
