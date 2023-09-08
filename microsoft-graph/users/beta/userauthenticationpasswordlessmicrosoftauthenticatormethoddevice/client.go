package userauthenticationpasswordlessmicrosoftauthenticatormethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationpasswordlessmicrosoftauthenticatormethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient: %+v", err)
	}

	return &UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient{
		Client: client,
	}, nil
}
