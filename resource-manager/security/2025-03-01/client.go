package v2025_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2025-03-01/devops"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DevOps *devops.DevOpsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	devOpsClient, err := devops.NewDevOpsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DevOps client: %+v", err)
	}
	configureFunc(devOpsClient.Client)

	return &Client{
		DevOps: devOpsClient,
	}, nil
}
