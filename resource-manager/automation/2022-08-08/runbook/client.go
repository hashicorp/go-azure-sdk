package runbook

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunbookClient struct {
	Client *resourcemanager.Client
}

func NewRunbookClientWithBaseURI(api environments.Api) (*RunbookClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "runbook", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RunbookClient: %+v", err)
	}

	return &RunbookClient{
		Client: client,
	}, nil
}
