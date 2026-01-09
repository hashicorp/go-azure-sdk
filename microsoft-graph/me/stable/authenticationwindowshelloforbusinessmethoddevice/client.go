package authenticationwindowshelloforbusinessmethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationWindowsHelloForBusinessMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationWindowsHelloForBusinessMethodDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationWindowsHelloForBusinessMethodDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationwindowshelloforbusinessmethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationWindowsHelloForBusinessMethodDeviceClient: %+v", err)
	}

	return &AuthenticationWindowsHelloForBusinessMethodDeviceClient{
		Client: client,
	}, nil
}
