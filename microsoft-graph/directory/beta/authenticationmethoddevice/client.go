package authenticationmethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMethodDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMethodDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMethodDeviceClient: %+v", err)
	}

	return &AuthenticationMethodDeviceClient{
		Client: client,
	}, nil
}
