package componentversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentVersionClient struct {
	Client *resourcemanager.Client
}

func NewComponentVersionClientWithBaseURI(api sdkEnv.Api) (*ComponentVersionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "componentversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComponentVersionClient: %+v", err)
	}

	return &ComponentVersionClient{
		Client: client,
	}, nil
}
