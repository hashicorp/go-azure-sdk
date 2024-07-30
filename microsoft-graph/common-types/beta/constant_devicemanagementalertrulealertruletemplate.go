package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRuleAlertRuleTemplate string

const (
	DeviceManagementAlertRuleAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario   DeviceManagementAlertRuleAlertRuleTemplate = "cloudPcFrontlineInsufficientLicensesScenario"
	DeviceManagementAlertRuleAlertRuleTemplate_CloudPCImageUploadScenario                     DeviceManagementAlertRuleAlertRuleTemplate = "cloudPcImageUploadScenario"
	DeviceManagementAlertRuleAlertRuleTemplate_CloudPCInGracePeriodScenario                   DeviceManagementAlertRuleAlertRuleTemplate = "cloudPcInGracePeriodScenario"
	DeviceManagementAlertRuleAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario DeviceManagementAlertRuleAlertRuleTemplate = "cloudPcOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementAlertRuleAlertRuleTemplate_CloudPCProvisionScenario                       DeviceManagementAlertRuleAlertRuleTemplate = "cloudPcProvisionScenario"
)

func PossibleValuesForDeviceManagementAlertRuleAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementAlertRuleAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplate_CloudPCImageUploadScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplate_CloudPCInGracePeriodScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementAlertRuleAlertRuleTemplate_CloudPCProvisionScenario),
	}
}

func (s *DeviceManagementAlertRuleAlertRuleTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertRuleAlertRuleTemplate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertRuleAlertRuleTemplate(input string) (*DeviceManagementAlertRuleAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementAlertRuleAlertRuleTemplate{
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementAlertRuleAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementAlertRuleAlertRuleTemplate_CloudPCImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementAlertRuleAlertRuleTemplate_CloudPCInGracePeriodScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementAlertRuleAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementAlertRuleAlertRuleTemplate_CloudPCProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRuleAlertRuleTemplate(input)
	return &out, nil
}
