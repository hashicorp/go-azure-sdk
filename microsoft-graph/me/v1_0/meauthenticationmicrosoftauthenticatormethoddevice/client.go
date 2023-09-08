package meauthenticationmicrosoftauthenticatormethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationMicrosoftAuthenticatorMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationMicrosoftAuthenticatorMethodDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationmicrosoftauthenticatormethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationMicrosoftAuthenticatorMethodDeviceClient: %+v", err)
	}

	return &MeAuthenticationMicrosoftAuthenticatorMethodDeviceClient{
		Client: client,
	}, nil
}
