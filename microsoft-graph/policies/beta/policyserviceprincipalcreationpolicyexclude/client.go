package policyserviceprincipalcreationpolicyexclude

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyServicePrincipalCreationPolicyExcludeClient struct {
	Client *msgraph.Client
}

func NewPolicyServicePrincipalCreationPolicyExcludeClientWithBaseURI(api sdkEnv.Api) (*PolicyServicePrincipalCreationPolicyExcludeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyserviceprincipalcreationpolicyexclude", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyServicePrincipalCreationPolicyExcludeClient: %+v", err)
	}

	return &PolicyServicePrincipalCreationPolicyExcludeClient{
		Client: client,
	}, nil
}
