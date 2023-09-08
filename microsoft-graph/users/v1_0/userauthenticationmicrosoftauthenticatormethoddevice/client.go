package userauthenticationmicrosoftauthenticatormethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationMicrosoftAuthenticatorMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationMicrosoftAuthenticatorMethodDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationmicrosoftauthenticatormethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationMicrosoftAuthenticatorMethodDeviceClient: %+v", err)
	}

	return &UserAuthenticationMicrosoftAuthenticatorMethodDeviceClient{
		Client: client,
	}, nil
}
