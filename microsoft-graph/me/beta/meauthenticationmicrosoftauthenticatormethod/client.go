package meauthenticationmicrosoftauthenticatormethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAuthenticationMicrosoftAuthenticatorMethodClient struct {
	Client *msgraph.Client
}

func NewMeAuthenticationMicrosoftAuthenticatorMethodClientWithBaseURI(api sdkEnv.Api) (*MeAuthenticationMicrosoftAuthenticatorMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meauthenticationmicrosoftauthenticatormethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAuthenticationMicrosoftAuthenticatorMethodClient: %+v", err)
	}

	return &MeAuthenticationMicrosoftAuthenticatorMethodClient{
		Client: client,
	}, nil
}
