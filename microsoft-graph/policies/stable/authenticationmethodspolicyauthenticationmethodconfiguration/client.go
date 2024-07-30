package authenticationmethodspolicyauthenticationmethodconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient struct {
	Client *msgraph.Client
}

func NewAuthenticationMethodsPolicyAuthenticationMethodConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "authenticationmethodspolicyauthenticationmethodconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient: %+v", err)
	}

	return &AuthenticationMethodsPolicyAuthenticationMethodConfigurationClient{
		Client: client,
	}, nil
}
