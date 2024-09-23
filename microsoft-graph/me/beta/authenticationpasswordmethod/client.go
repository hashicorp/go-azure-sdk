package authenticationpasswordmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationPasswordMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationPasswordMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationPasswordMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationpasswordmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationPasswordMethodClient: %+v", err)
	}

	return &AuthenticationPasswordMethodClient{
		Client: client,
	}, nil
}
