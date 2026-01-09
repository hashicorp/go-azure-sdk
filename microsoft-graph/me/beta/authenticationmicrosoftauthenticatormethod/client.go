package authenticationmicrosoftauthenticatormethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMicrosoftAuthenticatorMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMicrosoftAuthenticatorMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMicrosoftAuthenticatorMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmicrosoftauthenticatormethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMicrosoftAuthenticatorMethodClient: %+v", err)
	}

	return &AuthenticationMicrosoftAuthenticatorMethodClient{
		Client: client,
	}, nil
}
