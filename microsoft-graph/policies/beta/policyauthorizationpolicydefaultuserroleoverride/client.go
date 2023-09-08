package policyauthorizationpolicydefaultuserroleoverride

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyAuthorizationPolicyDefaultUserRoleOverrideClient struct {
	Client *msgraph.Client
}

func NewPolicyAuthorizationPolicyDefaultUserRoleOverrideClientWithBaseURI(api sdkEnv.Api) (*PolicyAuthorizationPolicyDefaultUserRoleOverrideClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyauthorizationpolicydefaultuserroleoverride", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyAuthorizationPolicyDefaultUserRoleOverrideClient: %+v", err)
	}

	return &PolicyAuthorizationPolicyDefaultUserRoleOverrideClient{
		Client: client,
	}, nil
}
