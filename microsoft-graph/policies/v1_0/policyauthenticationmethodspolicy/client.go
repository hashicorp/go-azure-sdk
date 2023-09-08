package policyauthenticationmethodspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAuthenticationMethodsPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyAuthenticationMethodsPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyAuthenticationMethodsPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyauthenticationmethodspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAuthenticationMethodsPolicyClient: %+v", err)
	}

	return &PolicyAuthenticationMethodsPolicyClient{
		Client: client,
	}, nil
}
