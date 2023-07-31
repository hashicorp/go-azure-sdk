package runs

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunsClient struct {
	Client *resourcemanager.Client
}

func NewRunsClientWithBaseURI(api sdkEnv.Api) (*RunsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "runs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RunsClient: %+v", err)
	}

	return &RunsClient{
		Client: client,
	}, nil
}
