package authenticationpasswordlessmicrosoftauthenticatormethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationPasswordlessMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationpasswordlessmicrosoftauthenticatormethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient: %+v", err)
	}

	return &AuthenticationPasswordlessMicrosoftAuthenticatorMethodClient{
		Client: client,
	}, nil
}
