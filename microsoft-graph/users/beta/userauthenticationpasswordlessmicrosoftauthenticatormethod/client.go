package userauthenticationpasswordlessmicrosoftauthenticatormethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient struct {
	Client *msgraph.Client
}

func NewUserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClientWithBaseURI(api sdkEnv.Api) (*UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userauthenticationpasswordlessmicrosoftauthenticatormethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient: %+v", err)
	}

	return &UserAuthenticationPasswordlessMicrosoftAuthenticatorMethodClient{
		Client: client,
	}, nil
}
