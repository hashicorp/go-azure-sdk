package v2023_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/resources/2023-08-01/deploymentscripts"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentScripts *deploymentscripts.DeploymentScriptsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deploymentScriptsClient, err := deploymentscripts.NewDeploymentScriptsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentScripts client: %+v", err)
	}
	configureFunc(deploymentScriptsClient.Client)

	return &Client{
		DeploymentScripts: deploymentScriptsClient,
	}, nil
}
