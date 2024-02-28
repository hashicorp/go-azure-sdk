package iotsecuritysolutions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotSecuritySolutionsClient struct {
	Client *resourcemanager.Client
}

func NewIotSecuritySolutionsClientWithBaseURI(sdkApi sdkEnv.Api) (*IotSecuritySolutionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "iotsecuritysolutions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IotSecuritySolutionsClient: %+v", err)
	}

	return &IotSecuritySolutionsClient{
		Client: client,
	}, nil
}
