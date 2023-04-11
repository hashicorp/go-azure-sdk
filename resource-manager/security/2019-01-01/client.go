package v2019_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/advancedthreatprotection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/alerts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2019-01-01/settings"
)

type Client struct {
	AdvancedThreatProtection *advancedthreatprotection.AdvancedThreatProtectionClient
	Alerts                   *alerts.AlertsClient
	Settings                 *settings.SettingsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	advancedThreatProtectionClient := advancedthreatprotection.NewAdvancedThreatProtectionClientWithBaseURI(endpoint)
	configureAuthFunc(&advancedThreatProtectionClient.Client)

	alertsClient := alerts.NewAlertsClientWithBaseURI(endpoint)
	configureAuthFunc(&alertsClient.Client)

	settingsClient := settings.NewSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&settingsClient.Client)

	return Client{
		AdvancedThreatProtection: &advancedThreatProtectionClient,
		Alerts:                   &alertsClient,
		Settings:                 &settingsClient,
	}
}
