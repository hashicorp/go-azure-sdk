package featurerolloutpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureRolloutPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewFeatureRolloutPolicyAppliesToClientWithBaseURI(sdkApi sdkEnv.Api) (*FeatureRolloutPolicyAppliesToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "featurerolloutpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeatureRolloutPolicyAppliesToClient: %+v", err)
	}

	return &FeatureRolloutPolicyAppliesToClient{
		Client: client,
	}, nil
}
