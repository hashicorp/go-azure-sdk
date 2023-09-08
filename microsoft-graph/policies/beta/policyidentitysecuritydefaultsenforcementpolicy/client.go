package policyidentitysecuritydefaultsenforcementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyIdentitySecurityDefaultsEnforcementPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyIdentitySecurityDefaultsEnforcementPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyIdentitySecurityDefaultsEnforcementPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyidentitysecuritydefaultsenforcementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyIdentitySecurityDefaultsEnforcementPolicyClient: %+v", err)
	}

	return &PolicyIdentitySecurityDefaultsEnforcementPolicyClient{
		Client: client,
	}, nil
}
