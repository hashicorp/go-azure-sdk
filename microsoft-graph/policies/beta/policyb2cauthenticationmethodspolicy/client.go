package policyb2cauthenticationmethodspolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyB2cAuthenticationMethodsPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyB2cAuthenticationMethodsPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyB2cAuthenticationMethodsPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyb2cauthenticationmethodspolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyB2cAuthenticationMethodsPolicyClient: %+v", err)
	}

	return &PolicyB2cAuthenticationMethodsPolicyClient{
		Client: client,
	}, nil
}
