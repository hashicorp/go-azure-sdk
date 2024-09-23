package authenticationwindowshelloforbusinessmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationWindowsHelloForBusinessMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationWindowsHelloForBusinessMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationWindowsHelloForBusinessMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationwindowshelloforbusinessmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationWindowsHelloForBusinessMethodClient: %+v", err)
	}

	return &AuthenticationWindowsHelloForBusinessMethodClient{
		Client: client,
	}, nil
}
