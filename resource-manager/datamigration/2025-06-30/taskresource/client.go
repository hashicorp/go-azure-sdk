package taskresource

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskResourceClient struct {
	Client *resourcemanager.Client
}

func NewTaskResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*TaskResourceClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "taskresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TaskResourceClient: %+v", err)
	}

	return &TaskResourceClient{
		Client: client,
	}, nil
}
