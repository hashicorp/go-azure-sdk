package jobtargetexecutions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobTargetExecutionsClient struct {
	Client *resourcemanager.Client
}

func NewJobTargetExecutionsClientWithBaseURI(sdkApi sdkEnv.Api) (*JobTargetExecutionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "jobtargetexecutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating JobTargetExecutionsClient: %+v", err)
	}

	return &JobTargetExecutionsClient{
		Client: client,
	}, nil
}
