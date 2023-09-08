package policyfeaturerolloutpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyFeatureRolloutPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyFeatureRolloutPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyFeatureRolloutPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyfeaturerolloutpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyFeatureRolloutPolicyClient: %+v", err)
	}

	return &PolicyFeatureRolloutPolicyClient{
		Client: client,
	}, nil
}
