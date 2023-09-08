package policyfeaturerolloutpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyFeatureRolloutPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyFeatureRolloutPolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyFeatureRolloutPolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyfeaturerolloutpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyFeatureRolloutPolicyAppliesToClient: %+v", err)
	}

	return &PolicyFeatureRolloutPolicyAppliesToClient{
		Client: client,
	}, nil
}
