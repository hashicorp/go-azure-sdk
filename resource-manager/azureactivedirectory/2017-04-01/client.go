package v2017_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2017-04-01/diagnosticsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azureactivedirectory/2017-04-01/diagnosticsettingscategories"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DiagnosticSettings           *diagnosticsettings.DiagnosticSettingsClient
	DiagnosticSettingsCategories *diagnosticsettingscategories.DiagnosticSettingsCategoriesClient
}

func NewClientWithBaseURI(api sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	diagnosticSettingsCategoriesClient, err := diagnosticsettingscategories.NewDiagnosticSettingsCategoriesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticSettingsCategories client: %+v", err)
	}
	configureFunc(diagnosticSettingsCategoriesClient.Client)

	diagnosticSettingsClient, err := diagnosticsettings.NewDiagnosticSettingsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DiagnosticSettings client: %+v", err)
	}
	configureFunc(diagnosticSettingsClient.Client)

	return &Client{
		DiagnosticSettings:           diagnosticSettingsClient,
		DiagnosticSettingsCategories: diagnosticSettingsCategoriesClient,
	}, nil
}
