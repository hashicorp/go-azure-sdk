package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRuleAlertRuleTemplate string

const (
	DeviceManagementAlertRuleAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario   DeviceManagementAlertRuleAlertRuleTemplate = "CloudPCFrontlineInsufficientLicensesScenario"
	DeviceManagementAlertRuleAlertRuleTemplatecloudPcImageUploadScenario                     DeviceManagementAlertRuleAlertRuleTemplate = "CloudPCImageUploadScenario"
	DeviceManagementAlertRuleAlertRuleTemplatecloudPcInGracePeriodScenario                   DeviceManagementAlertRuleAlertRuleTemplate = "CloudPCInGracePeriodScenario"
	DeviceManagementAlertRuleAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario DeviceManagementAlertRuleAlertRuleTemplate = "CloudPCOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementAlertRuleAlertRuleTemplatecloudPcProvisionScenario                       DeviceManagementAlertRuleAlertRuleTemplate = "CloudPCProvisionScenario"
)

func PossibleValuesForDeviceManagementAlertRuleAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementAlertRuleAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplatecloudPcImageUploadScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplatecloudPcInGracePeriodScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplatecloudPcProvisionScenario),
	}
}

func parseDeviceManagementAlertRuleAlertRuleTemplate(input string) (*DeviceManagementAlertRuleAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementAlertRuleAlertRuleTemplate{
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementAlertRuleAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementAlertRuleAlertRuleTemplatecloudPcImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementAlertRuleAlertRuleTemplatecloudPcInGracePeriodScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementAlertRuleAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementAlertRuleAlertRuleTemplatecloudPcProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRuleAlertRuleTemplate(input)
	return &out, nil
}
