package meauthenticationpasswordlessmicrosoftauthenticatormethoddevice

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationpasswordlessmicrosoftauthenticatormethoddevice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient: %+v", err)
	}

	return &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodDeviceClient{
		Client: client,
	}, nil
}
