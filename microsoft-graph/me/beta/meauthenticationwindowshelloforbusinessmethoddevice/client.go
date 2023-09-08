package meauthenticationwindowshelloforbusinessmethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationWindowsHelloForBusinessMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationWindowsHelloForBusinessMethodDeviceClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationWindowsHelloForBusinessMethodDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationwindowshelloforbusinessmethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationWindowsHelloForBusinessMethodDeviceClient: %+v", err)
	}

	return &MeAuthenticationWindowsHelloForBusinessMethodDeviceClient{
		Client: client,
	}, nil
}
