package privateaccesses

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateAccessesClient struct {
	Client *resourcemanager.Client
}

func NewPrivateAccessesClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateAccessesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "privateaccesses", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateAccessesClient: %+v", err)
	}

	return &PrivateAccessesClient{
		Client: client,
	}, nil
}
