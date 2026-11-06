package daprcomponentresiliencypolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DaprComponentResiliencyPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewDaprComponentResiliencyPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*DaprComponentResiliencyPoliciesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "daprcomponentresiliencypolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DaprComponentResiliencyPoliciesClient: %+v", err)
	}

	return &DaprComponentResiliencyPoliciesClient{
		Client: client,
	}, nil
}
