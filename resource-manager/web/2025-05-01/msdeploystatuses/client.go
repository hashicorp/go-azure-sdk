package msdeploystatuses

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MSDeployStatusesClient struct {
	Client *resourcemanager.Client
}

func NewMSDeployStatusesClientWithBaseURI(sdkApi sdkEnv.Api) (*MSDeployStatusesClient, error) {
	client, err := resourcemanager.NewClient(sdkApi, "msdeploystatuses", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MSDeployStatusesClient: %+v", err)
	}

	return &MSDeployStatusesClient{
		Client: client,
	}, nil
}
