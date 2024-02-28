package v2019_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/advancedthreatprotection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/settings"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AdvancedThreatProtection *advancedthreatprotection.AdvancedThreatProtectionClient
	Alerts                   *alerts.AlertsClient
	Settings                 *settings.SettingsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	advancedThreatProtectionClient, err := advancedthreatprotection.NewAdvancedThreatProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdvancedThreatProtection client: %+v", err)
	}
	configureFunc(advancedThreatProtectionClient.Client)

	alertsClient, err := alerts.NewAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	settingsClient, err := settings.NewSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %+v", err)
	}
	configureFunc(settingsClient.Client)

	return &Client{
		AdvancedThreatProtection: advancedThreatProtectionClient,
		Alerts:                   alertsClient,
		Settings:                 settingsClient,
	}, nil
}
