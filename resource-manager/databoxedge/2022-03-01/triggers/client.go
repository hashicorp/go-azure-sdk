package triggers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggersClient struct {
	Client *resourcemanager.Client
}

func NewTriggersClientWithBaseURI(api sdkEnv.Api) (*TriggersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "triggers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggersClient: %+v", err)
	}

	return &TriggersClient{
		Client: client,
	}, nil
}
