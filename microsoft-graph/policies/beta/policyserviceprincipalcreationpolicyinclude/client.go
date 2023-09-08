package policyserviceprincipalcreationpolicyinclude

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyServicePrincipalCreationPolicyIncludeClient struct {
	Client *msgraph.Client
}

func NewPolicyServicePrincipalCreationPolicyIncludeClientWithBaseURI(api sdkEnv.Api) (*PolicyServicePrincipalCreationPolicyIncludeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyserviceprincipalcreationpolicyinclude", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyServicePrincipalCreationPolicyIncludeClient: %+v", err)
	}

	return &PolicyServicePrincipalCreationPolicyIncludeClient{
		Client: client,
	}, nil
}
