package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPortalNotificationAlertRuleTemplate string

const (
	DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario   DeviceManagementPortalNotificationAlertRuleTemplate = "cloudPcFrontlineInsufficientLicensesScenario"
	DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCImageUploadScenario                     DeviceManagementPortalNotificationAlertRuleTemplate = "cloudPcImageUploadScenario"
	DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCInGracePeriodScenario                   DeviceManagementPortalNotificationAlertRuleTemplate = "cloudPcInGracePeriodScenario"
	DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario DeviceManagementPortalNotificationAlertRuleTemplate = "cloudPcOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCProvisionScenario                       DeviceManagementPortalNotificationAlertRuleTemplate = "cloudPcProvisionScenario"
)

func PossibleValuesForDeviceManagementPortalNotificationAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCImageUploadScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCInGracePeriodScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCProvisionScenario),
	}
}

func (s *DeviceManagementPortalNotificationAlertRuleTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementPortalNotificationAlertRuleTemplate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementPortalNotificationAlertRuleTemplate(input string) (*DeviceManagementPortalNotificationAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementPortalNotificationAlertRuleTemplate{
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCInGracePeriodScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementPortalNotificationAlertRuleTemplate_CloudPCProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementPortalNotificationAlertRuleTemplate(input)
	return &out, nil
}
