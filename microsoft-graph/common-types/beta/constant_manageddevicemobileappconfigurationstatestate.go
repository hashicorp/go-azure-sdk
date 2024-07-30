package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationStateState string

const (
	ManagedDeviceMobileAppConfigurationStateState_Compliant     ManagedDeviceMobileAppConfigurationStateState = "compliant"
	ManagedDeviceMobileAppConfigurationStateState_Conflict      ManagedDeviceMobileAppConfigurationStateState = "conflict"
	ManagedDeviceMobileAppConfigurationStateState_Error         ManagedDeviceMobileAppConfigurationStateState = "error"
	ManagedDeviceMobileAppConfigurationStateState_NonCompliant  ManagedDeviceMobileAppConfigurationStateState = "nonCompliant"
	ManagedDeviceMobileAppConfigurationStateState_NotApplicable ManagedDeviceMobileAppConfigurationStateState = "notApplicable"
	ManagedDeviceMobileAppConfigurationStateState_NotAssigned   ManagedDeviceMobileAppConfigurationStateState = "notAssigned"
	ManagedDeviceMobileAppConfigurationStateState_Remediated    ManagedDeviceMobileAppConfigurationStateState = "remediated"
	ManagedDeviceMobileAppConfigurationStateState_Unknown       ManagedDeviceMobileAppConfigurationStateState = "unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationStateState() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationStateState_Compliant),
		string(ManagedDeviceMobileAppConfigurationStateState_Conflict),
		string(ManagedDeviceMobileAppConfigurationStateState_Error),
		string(ManagedDeviceMobileAppConfigurationStateState_NonCompliant),
		string(ManagedDeviceMobileAppConfigurationStateState_NotApplicable),
		string(ManagedDeviceMobileAppConfigurationStateState_NotAssigned),
		string(ManagedDeviceMobileAppConfigurationStateState_Remediated),
		string(ManagedDeviceMobileAppConfigurationStateState_Unknown),
	}
}

func (s *ManagedDeviceMobileAppConfigurationStateState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationStateState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationStateState(input string) (*ManagedDeviceMobileAppConfigurationStateState, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationStateState{
		"compliant":     ManagedDeviceMobileAppConfigurationStateState_Compliant,
		"conflict":      ManagedDeviceMobileAppConfigurationStateState_Conflict,
		"error":         ManagedDeviceMobileAppConfigurationStateState_Error,
		"noncompliant":  ManagedDeviceMobileAppConfigurationStateState_NonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationStateState_NotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationStateState_NotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationStateState_Remediated,
		"unknown":       ManagedDeviceMobileAppConfigurationStateState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationStateState(input)
	return &out, nil
}
