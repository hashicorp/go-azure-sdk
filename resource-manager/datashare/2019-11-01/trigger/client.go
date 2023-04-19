package trigger

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerClient struct {
	Client *resourcemanager.Client
}

func NewTriggerClientWithBaseURI(api environments.Api) (*TriggerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "trigger", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggerClient: %+v", err)
	}

	return &TriggerClient{
		Client: client,
	}, nil
}
