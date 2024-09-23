package identitysecuritydefaultsenforcementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentitySecurityDefaultsEnforcementPolicyClient struct {
	Client *msgraph.Client
}

func NewIdentitySecurityDefaultsEnforcementPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*IdentitySecurityDefaultsEnforcementPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "identitysecuritydefaultsenforcementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentitySecurityDefaultsEnforcementPolicyClient: %+v", err)
	}

	return &IdentitySecurityDefaultsEnforcementPolicyClient{
		Client: client,
	}, nil
}
