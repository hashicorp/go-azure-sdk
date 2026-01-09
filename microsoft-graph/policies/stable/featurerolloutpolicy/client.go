package featurerolloutpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureRolloutPolicyClient struct {
	Client *msgraph.Client
}

func NewFeatureRolloutPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*FeatureRolloutPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "featurerolloutpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeatureRolloutPolicyClient: %+v", err)
	}

	return &FeatureRolloutPolicyClient{
		Client: client,
	}, nil
}
