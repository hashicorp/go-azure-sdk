package v2025_09_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2025-09-01/playwrightquotas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2025-09-01/playwrightworkspacequotas"
	"github.com/hashicorp/go-azure-sdk/resource-manager/loadtestservice/2025-09-01/playwrightworkspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	PlaywrightQuotas          *playwrightquotas.PlaywrightQuotasClient
	PlaywrightWorkspaceQuotas *playwrightworkspacequotas.PlaywrightWorkspaceQuotasClient
	PlaywrightWorkspaces      *playwrightworkspaces.PlaywrightWorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	playwrightQuotasClient, err := playwrightquotas.NewPlaywrightQuotasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlaywrightQuotas client: %+v", err)
	}
	configureFunc(playwrightQuotasClient.Client)

	playwrightWorkspaceQuotasClient, err := playwrightworkspacequotas.NewPlaywrightWorkspaceQuotasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlaywrightWorkspaceQuotas client: %+v", err)
	}
	configureFunc(playwrightWorkspaceQuotasClient.Client)

	playwrightWorkspacesClient, err := playwrightworkspaces.NewPlaywrightWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PlaywrightWorkspaces client: %+v", err)
	}
	configureFunc(playwrightWorkspacesClient.Client)

	return &Client{
		PlaywrightQuotas:          playwrightQuotasClient,
		PlaywrightWorkspaceQuotas: playwrightWorkspaceQuotasClient,
		PlaywrightWorkspaces:      playwrightWorkspacesClient,
	}, nil
}
