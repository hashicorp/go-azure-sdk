package componentworkitemconfigsapis

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComponentWorkItemConfigsAPIsClient struct {
	Client *resourcemanager.Client
}

func NewComponentWorkItemConfigsAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*ComponentWorkItemConfigsAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "componentworkitemconfigsapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComponentWorkItemConfigsAPIsClient: %+v", err)
	}

	return &ComponentWorkItemConfigsAPIsClient{
		Client: client,
	}, nil
}
