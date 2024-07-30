package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationUserStatusStatus string

const (
	ManagedDeviceMobileAppConfigurationUserStatusStatus_Compliant     ManagedDeviceMobileAppConfigurationUserStatusStatus = "compliant"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_Conflict      ManagedDeviceMobileAppConfigurationUserStatusStatus = "conflict"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_Error         ManagedDeviceMobileAppConfigurationUserStatusStatus = "error"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_NonCompliant  ManagedDeviceMobileAppConfigurationUserStatusStatus = "nonCompliant"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_NotApplicable ManagedDeviceMobileAppConfigurationUserStatusStatus = "notApplicable"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_NotAssigned   ManagedDeviceMobileAppConfigurationUserStatusStatus = "notAssigned"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_Remediated    ManagedDeviceMobileAppConfigurationUserStatusStatus = "remediated"
	ManagedDeviceMobileAppConfigurationUserStatusStatus_Unknown       ManagedDeviceMobileAppConfigurationUserStatusStatus = "unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationUserStatusStatus() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_Compliant),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_Conflict),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_Error),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_NonCompliant),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_NotApplicable),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_NotAssigned),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_Remediated),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatus_Unknown),
	}
}

func (s *ManagedDeviceMobileAppConfigurationUserStatusStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceMobileAppConfigurationUserStatusStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceMobileAppConfigurationUserStatusStatus(input string) (*ManagedDeviceMobileAppConfigurationUserStatusStatus, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationUserStatusStatus{
		"compliant":     ManagedDeviceMobileAppConfigurationUserStatusStatus_Compliant,
		"conflict":      ManagedDeviceMobileAppConfigurationUserStatusStatus_Conflict,
		"error":         ManagedDeviceMobileAppConfigurationUserStatusStatus_Error,
		"noncompliant":  ManagedDeviceMobileAppConfigurationUserStatusStatus_NonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationUserStatusStatus_NotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationUserStatusStatus_NotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationUserStatusStatus_Remediated,
		"unknown":       ManagedDeviceMobileAppConfigurationUserStatusStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationUserStatusStatus(input)
	return &out, nil
}
