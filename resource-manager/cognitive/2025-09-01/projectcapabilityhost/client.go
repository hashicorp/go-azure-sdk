package projectcapabilityhost

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectCapabilityHostClient struct {
	Client *resourcemanager.Client
}

func NewProjectCapabilityHostClientWithBaseURI(sdkApi sdkEnv.Api) (*ProjectCapabilityHostClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "projectcapabilityhost", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ProjectCapabilityHostClient: %+v", err)
	}

	return &ProjectCapabilityHostClient{
		Client: client,
	}, nil
}
