package jobcredentials

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobCredentialsClient struct {
	Client *resourcemanager.Client
}

func NewJobCredentialsClientWithBaseURI(api sdkEnv.Api) (*JobCredentialsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "jobcredentials", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobCredentialsClient: %+v", err)
	}

	return &JobCredentialsClient{
		Client: client,
	}, nil
}
