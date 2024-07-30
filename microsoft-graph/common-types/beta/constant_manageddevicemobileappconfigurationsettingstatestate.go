package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationSettingStateState string

const (
	ManagedDeviceMobileAppConfigurationSettingStateState_Compliant     ManagedDeviceMobileAppConfigurationSettingStateState = "compliant"
	ManagedDeviceMobileAppConfigurationSettingStateState_Conflict      ManagedDeviceMobileAppConfigurationSettingStateState = "conflict"
	ManagedDeviceMobileAppConfigurationSettingStateState_Error         ManagedDeviceMobileAppConfigurationSettingStateState = "error"
	ManagedDeviceMobileAppConfigurationSettingStateState_NonCompliant  ManagedDeviceMobileAppConfigurationSettingStateState = "nonCompliant"
	ManagedDeviceMobileAppConfigurationSettingStateState_NotApplicable ManagedDeviceMobileAppConfigurationSettingStateState = "notApplicable"
	ManagedDeviceMobileAppConfigurationSettingStateState_NotAssigned   ManagedDeviceMobileAppConfigurationSettingStateState = "notAssigned"
	ManagedDeviceMobileAppConfigurationSettingStateState_Remediated    ManagedDeviceMobileAppConfigurationSettingStateState = "remediated"
	ManagedDeviceMobileAppConfigurationSettingStateState_Unknown       ManagedDeviceMobileAppConfigurationSettingStateState = "unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationSettingStateState() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationSettingStateState_Compliant),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_Conflict),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_Error),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_NonCompliant),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_NotApplicable),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_NotAssigned),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_Remediated),
		string(ManagedDeviceMobileAppConfigurationSettingStateState_Unknown),
	}
}

func (s *ManagedDeviceMobileAppConfigurationSettingStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationSettingStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationSettingStateState(input string) (*ManagedDeviceMobileAppConfigurationSettingStateState, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationSettingStateState{
		"compliant":     ManagedDeviceMobileAppConfigurationSettingStateState_Compliant,
		"conflict":      ManagedDeviceMobileAppConfigurationSettingStateState_Conflict,
		"error":         ManagedDeviceMobileAppConfigurationSettingStateState_Error,
		"noncompliant":  ManagedDeviceMobileAppConfigurationSettingStateState_NonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationSettingStateState_NotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationSettingStateState_NotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationSettingStateState_Remediated,
		"unknown":       ManagedDeviceMobileAppConfigurationSettingStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationSettingStateState(input)
	return &out, nil
}
