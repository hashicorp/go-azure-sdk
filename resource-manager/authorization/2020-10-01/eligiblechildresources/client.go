package eligiblechildresources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EligibleChildResourcesClient struct {
	Client *resourcemanager.Client
}

func NewEligibleChildResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*EligibleChildResourcesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "eligiblechildresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EligibleChildResourcesClient: %+v", err)
	}

	return &EligibleChildResourcesClient{
		Client: client,
	}, nil
}
