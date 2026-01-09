package authenticationeventlistener

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEventListenerClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEventListenerClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEventListenerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationeventlistener", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEventListenerClient: %+v", err)
	}

	return &AuthenticationEventListenerClient{
		Client: client,
	}, nil
}
