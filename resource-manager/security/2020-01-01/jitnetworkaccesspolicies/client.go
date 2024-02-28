package jitnetworkaccesspolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitNetworkAccessPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewJitNetworkAccessPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*JitNetworkAccessPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "jitnetworkaccesspolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JitNetworkAccessPoliciesClient: %+v", err)
	}

	return &JitNetworkAccessPoliciesClient{
		Client: client,
	}, nil
}
