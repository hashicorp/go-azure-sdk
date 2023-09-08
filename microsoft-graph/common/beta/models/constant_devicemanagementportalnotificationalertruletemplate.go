package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPortalNotificationAlertRuleTemplate string

const (
	DeviceManagementPortalNotificationAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario   DeviceManagementPortalNotificationAlertRuleTemplate = "CloudPCFrontlineInsufficientLicensesScenario"
	DeviceManagementPortalNotificationAlertRuleTemplatecloudPcImageUploadScenario                     DeviceManagementPortalNotificationAlertRuleTemplate = "CloudPCImageUploadScenario"
	DeviceManagementPortalNotificationAlertRuleTemplatecloudPcInGracePeriodScenario                   DeviceManagementPortalNotificationAlertRuleTemplate = "CloudPCInGracePeriodScenario"
	DeviceManagementPortalNotificationAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario DeviceManagementPortalNotificationAlertRuleTemplate = "CloudPCOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementPortalNotificationAlertRuleTemplatecloudPcProvisionScenario                       DeviceManagementPortalNotificationAlertRuleTemplate = "CloudPCProvisionScenario"
)

func PossibleValuesForDeviceManagementPortalNotificationAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementPortalNotificationAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplatecloudPcImageUploadScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplatecloudPcInGracePeriodScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplatecloudPcProvisionScenario),
	}
}

func parseDeviceManagementPortalNotificationAlertRuleTemplate(input string) (*DeviceManagementPortalNotificationAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementPortalNotificationAlertRuleTemplate{
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementPortalNotificationAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementPortalNotificationAlertRuleTemplatecloudPcImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementPortalNotificationAlertRuleTemplatecloudPcInGracePeriodScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementPortalNotificationAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementPortalNotificationAlertRuleTemplatecloudPcProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPortalNotificationAlertRuleTemplate(input)
	return &out, nil
}
