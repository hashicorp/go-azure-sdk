package v2017_08_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

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
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	advancedThreatProtectionClient, err := advancedthreatprotection.NewAdvancedThreatProtectionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AdvancedThreatProtection client: %+v", err)
	}
	configureFunc(advancedThreatProtectionClient.Client)

	autoProvisioningSettingsClient, err := autoprovisioningsettings.NewAutoProvisioningSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AutoProvisioningSettings client: %+v", err)
	}
	configureFunc(autoProvisioningSettingsClient.Client)

	compliancesClient, err := compliances.NewCompliancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Compliances client: %+v", err)
	}
	configureFunc(compliancesClient.Client)

	deviceSecurityGroupsClient, err := devicesecuritygroups.NewDeviceSecurityGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeviceSecurityGroups client: %+v", err)
	}
	configureFunc(deviceSecurityGroupsClient.Client)

	informationProtectionPoliciesClient, err := informationprotectionpolicies.NewInformationProtectionPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InformationProtectionPolicies client: %+v", err)
	}
	configureFunc(informationProtectionPoliciesClient.Client)

	ioTSecuritySolutionsAnalyticsClient, err := iotsecuritysolutionsanalytics.NewIoTSecuritySolutionsAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IoTSecuritySolutionsAnalytics client: %+v", err)
	}
	configureFunc(ioTSecuritySolutionsAnalyticsClient.Client)

	iotSecuritySolutionsClient, err := iotsecuritysolutions.NewIotSecuritySolutionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IotSecuritySolutions client: %+v", err)
	}
	configureFunc(iotSecuritySolutionsClient.Client)

	pricingsClient, err := pricings.NewPricingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Pricings client: %+v", err)
	}
	configureFunc(pricingsClient.Client)

	securityContactsClient, err := securitycontacts.NewSecurityContactsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SecurityContacts client: %+v", err)
	}
	configureFunc(securityContactsClient.Client)

	settingsClient, err := settings.NewSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Settings client: %+v", err)
	}
	configureFunc(settingsClient.Client)

	workspaceSettingsClient, err := workspacesettings.NewWorkspaceSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkspaceSettings client: %+v", err)
	}
	configureFunc(workspaceSettingsClient.Client)

	return &Client{
		AdvancedThreatProtection:      advancedThreatProtectionClient,
		AutoProvisioningSettings:      autoProvisioningSettingsClient,
		Compliances:                   compliancesClient,
		DeviceSecurityGroups:          deviceSecurityGroupsClient,
		InformationProtectionPolicies: informationProtectionPoliciesClient,
		IoTSecuritySolutionsAnalytics: ioTSecuritySolutionsAnalyticsClient,
		IotSecuritySolutions:          iotSecuritySolutionsClient,
		Pricings:                      pricingsClient,
		SecurityContacts:              securityContactsClient,
		Settings:                      settingsClient,
		WorkspaceSettings:             workspaceSettingsClient,
	}, nil
}
