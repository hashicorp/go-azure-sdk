package authenticationfido2method

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationFido2MethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationFido2MethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationFido2MethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationfido2method", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationFido2MethodClient: %+v", err)
	}

	return &AuthenticationFido2MethodClient{
		Client: client,
	}, nil
}
