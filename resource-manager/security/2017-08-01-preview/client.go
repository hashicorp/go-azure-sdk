package v2017_08_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/advancedthreatprotection"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/autoprovisioningsettings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/compliances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/devicesecuritygroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/informationprotectionpolicies"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/iotsecuritysolutions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/iotsecuritysolutionsanalytics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/pricings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/securitycontacts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/settings"
	"github.com/hashicorp/go-azure-sdk/resource-manager/security/2017-08-01-preview/workspacesettings"
)

type Client struct {
	AdvancedThreatProtection      *advancedthreatprotection.AdvancedThreatProtectionClient
	AutoProvisioningSettings      *autoprovisioningsettings.AutoProvisioningSettingsClient
	Compliances                   *compliances.CompliancesClient
	DeviceSecurityGroups          *devicesecuritygroups.DeviceSecurityGroupsClient
	InformationProtectionPolicies *informationprotectionpolicies.InformationProtectionPoliciesClient
	IoTSecuritySolutionsAnalytics *iotsecuritysolutionsanalytics.IoTSecuritySolutionsAnalyticsClient
	IotSecuritySolutions          *iotsecuritysolutions.IotSecuritySolutionsClient
	Pricings                      *pricings.PricingsClient
	SecurityContacts              *securitycontacts.SecurityContactsClient
	Settings                      *settings.SettingsClient
	WorkspaceSettings             *workspacesettings.WorkspaceSettingsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	advancedThreatProtectionClient := advancedthreatprotection.NewAdvancedThreatProtectionClientWithBaseURI(endpoint)
	configureAuthFunc(&advancedThreatProtectionClient.Client)

	autoProvisioningSettingsClient := autoprovisioningsettings.NewAutoProvisioningSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&autoProvisioningSettingsClient.Client)

	compliancesClient := compliances.NewCompliancesClientWithBaseURI(endpoint)
	configureAuthFunc(&compliancesClient.Client)

	deviceSecurityGroupsClient := devicesecuritygroups.NewDeviceSecurityGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&deviceSecurityGroupsClient.Client)

	informationProtectionPoliciesClient := informationprotectionpolicies.NewInformationProtectionPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&informationProtectionPoliciesClient.Client)

	ioTSecuritySolutionsAnalyticsClient := iotsecuritysolutionsanalytics.NewIoTSecuritySolutionsAnalyticsClientWithBaseURI(endpoint)
	configureAuthFunc(&ioTSecuritySolutionsAnalyticsClient.Client)

	iotSecuritySolutionsClient := iotsecuritysolutions.NewIotSecuritySolutionsClientWithBaseURI(endpoint)
	configureAuthFunc(&iotSecuritySolutionsClient.Client)

	pricingsClient := pricings.NewPricingsClientWithBaseURI(endpoint)
	configureAuthFunc(&pricingsClient.Client)

	securityContactsClient := securitycontacts.NewSecurityContactsClientWithBaseURI(endpoint)
	configureAuthFunc(&securityContactsClient.Client)

	settingsClient := settings.NewSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&settingsClient.Client)

	workspaceSettingsClient := workspacesettings.NewWorkspaceSettingsClientWithBaseURI(endpoint)
	configureAuthFunc(&workspaceSettingsClient.Client)

	return Client{
		AdvancedThreatProtection:      &advancedThreatProtectionClient,
		AutoProvisioningSettings:      &autoProvisioningSettingsClient,
		Compliances:                   &compliancesClient,
		DeviceSecurityGroups:          &deviceSecurityGroupsClient,
		InformationProtectionPolicies: &informationProtectionPoliciesClient,
		IoTSecuritySolutionsAnalytics: &ioTSecuritySolutionsAnalyticsClient,
		IotSecuritySolutions:          &iotSecuritySolutionsClient,
		Pricings:                      &pricingsClient,
		SecurityContacts:              &securityContactsClient,
		Settings:                      &settingsClient,
		WorkspaceSettings:             &workspaceSettingsClient,
	}
}
