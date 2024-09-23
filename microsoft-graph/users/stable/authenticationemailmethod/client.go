package authenticationemailmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationEmailMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationEmailMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationEmailMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationemailmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationEmailMethodClient: %+v", err)
	}

	return &AuthenticationEmailMethodClient{
		Client: client,
	}, nil
}
