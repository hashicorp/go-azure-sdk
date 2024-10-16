package policymetadata

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyMetadataClient struct {
	Client *resourcemanager.Client
}

func NewPolicyMetadataClientWithBaseURI(sdkApi sdkEnv.Api) (*PolicyMetadataClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "policymetadata", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyMetadataClient: %+v", err)
	}

	return &PolicyMetadataClient{
		Client: client,
	}, nil
}
