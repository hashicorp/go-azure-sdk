package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRecordAlertRuleTemplate string

const (
	DeviceManagementAlertRecordAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario   DeviceManagementAlertRecordAlertRuleTemplate = "cloudPcFrontlineInsufficientLicensesScenario"
	DeviceManagementAlertRecordAlertRuleTemplate_CloudPCImageUploadScenario                     DeviceManagementAlertRecordAlertRuleTemplate = "cloudPcImageUploadScenario"
	DeviceManagementAlertRecordAlertRuleTemplate_CloudPCInGracePeriodScenario                   DeviceManagementAlertRecordAlertRuleTemplate = "cloudPcInGracePeriodScenario"
	DeviceManagementAlertRecordAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario DeviceManagementAlertRecordAlertRuleTemplate = "cloudPcOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementAlertRecordAlertRuleTemplate_CloudPCProvisionScenario                       DeviceManagementAlertRecordAlertRuleTemplate = "cloudPcProvisionScenario"
)

func PossibleValuesForDeviceManagementAlertRecordAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementAlertRecordAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplate_CloudPCImageUploadScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplate_CloudPCInGracePeriodScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementAlertRecordAlertRuleTemplate_CloudPCProvisionScenario),
	}
}

func (s *DeviceManagementAlertRecordAlertRuleTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertRecordAlertRuleTemplate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertRecordAlertRuleTemplate(input string) (*DeviceManagementAlertRecordAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementAlertRecordAlertRuleTemplate{
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementAlertRecordAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementAlertRecordAlertRuleTemplate_CloudPCImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementAlertRecordAlertRuleTemplate_CloudPCInGracePeriodScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementAlertRecordAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementAlertRecordAlertRuleTemplate_CloudPCProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRecordAlertRuleTemplate(input)
	return &out, nil
}
