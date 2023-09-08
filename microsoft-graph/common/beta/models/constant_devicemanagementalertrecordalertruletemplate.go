package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecordAlertRuleTemplate string

const (
	DeviceManagementAlertRecordAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario   DeviceManagementAlertRecordAlertRuleTemplate = "CloudPCFrontlineInsufficientLicensesScenario"
	DeviceManagementAlertRecordAlertRuleTemplatecloudPcImageUploadScenario                     DeviceManagementAlertRecordAlertRuleTemplate = "CloudPCImageUploadScenario"
	DeviceManagementAlertRecordAlertRuleTemplatecloudPcInGracePeriodScenario                   DeviceManagementAlertRecordAlertRuleTemplate = "CloudPCInGracePeriodScenario"
	DeviceManagementAlertRecordAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario DeviceManagementAlertRecordAlertRuleTemplate = "CloudPCOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementAlertRecordAlertRuleTemplatecloudPcProvisionScenario                       DeviceManagementAlertRecordAlertRuleTemplate = "CloudPCProvisionScenario"
)

func PossibleValuesForDeviceManagementAlertRecordAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementAlertRecordAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplatecloudPcImageUploadScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplatecloudPcInGracePeriodScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplatecloudPcProvisionScenario),
	}
}

func parseDeviceManagementAlertRecordAlertRuleTemplate(input string) (*DeviceManagementAlertRecordAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementAlertRecordAlertRuleTemplate{
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementAlertRecordAlertRuleTemplatecloudPcFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementAlertRecordAlertRuleTemplatecloudPcImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementAlertRecordAlertRuleTemplatecloudPcInGracePeriodScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementAlertRecordAlertRuleTemplatecloudPcOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementAlertRecordAlertRuleTemplatecloudPcProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRecordAlertRuleTemplate(input)
	return &out, nil
}
