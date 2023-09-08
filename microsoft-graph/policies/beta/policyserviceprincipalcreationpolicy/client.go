package policyserviceprincipalcreationpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyServicePrincipalCreationPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyServicePrincipalCreationPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyServicePrincipalCreationPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyserviceprincipalcreationpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyServicePrincipalCreationPolicyClient: %+v", err)
	}

	return &PolicyServicePrincipalCreationPolicyClient{
		Client: client,
	}, nil
}
