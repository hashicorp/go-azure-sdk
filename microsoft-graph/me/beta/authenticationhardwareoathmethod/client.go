package authenticationhardwareoathmethod

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationHardwareOathMethodClient struct {
	Client *msgraph.Client
}

func NewAuthenticationHardwareOathMethodClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationHardwareOathMethodClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationhardwareoathmethod", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationHardwareOathMethodClient: %+v", err)
	}

	return &AuthenticationHardwareOathMethodClient{
		Client: client,
	}, nil
}
