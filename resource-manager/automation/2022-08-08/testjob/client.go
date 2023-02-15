package testjob

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TestJobClient struct {
	Client *resourcemanager.Client
}

func NewTestJobClientWithBaseURI(api environments.Api) (*TestJobClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "testjob", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TestJobClient: %+v", err)
	}

	return &TestJobClient{
		Client: client,
	}, nil
}
