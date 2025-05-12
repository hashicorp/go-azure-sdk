package purestoragepolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PureStoragePoliciesClient struct {
	Client *resourcemanager.Client
}

func NewPureStoragePoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*PureStoragePoliciesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "purestoragepolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PureStoragePoliciesClient: %+v", err)
	}

	return &PureStoragePoliciesClient{
		Client: client,
	}, nil
}
