package authenticationmicrosoftauthenticatormethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMicrosoftAuthenticatorMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMicrosoftAuthenticatorMethodDeviceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmicrosoftauthenticatormethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMicrosoftAuthenticatorMethodDeviceClient: %+v", err)
	}

	return &AuthenticationMicrosoftAuthenticatorMethodDeviceClient{
		Client: client,
	}, nil
}
