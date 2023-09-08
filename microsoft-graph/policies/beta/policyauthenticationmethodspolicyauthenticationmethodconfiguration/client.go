package policyauthenticationmethodspolicyauthenticationmethodconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient struct {
	Client *msgraph.Client
}

func NewPolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClientWithBaseURI(api sdkEnv.Api) (*PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyauthenticationmethodspolicyauthenticationmethodconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient: %+v", err)
	}

	return &PolicyAuthenticationMethodsPolicyAuthenticationMethodConfigurationClient{
		Client: client,
	}, nil
}
