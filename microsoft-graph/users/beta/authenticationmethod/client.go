package authenticationmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMethodClient: %+v", err)
	}

	return &AuthenticationMethodClient{
		Client: client,
	}, nil
}
