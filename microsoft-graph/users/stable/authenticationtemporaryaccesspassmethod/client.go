package authenticationtemporaryaccesspassmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationTemporaryAccessPassMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationTemporaryAccessPassMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationTemporaryAccessPassMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationtemporaryaccesspassmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationTemporaryAccessPassMethodClient: %+v", err)
	}

	return &AuthenticationTemporaryAccessPassMethodClient{
		Client: client,
	}, nil
}
