package meauthenticationpasswordlessmicrosoftauthenticatormethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationpasswordlessmicrosoftauthenticatormethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient: %+v", err)
	}

	return &MeAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient{
		Client: client,
	}, nil
}
