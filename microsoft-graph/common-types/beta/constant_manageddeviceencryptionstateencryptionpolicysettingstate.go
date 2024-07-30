package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateEncryptionPolicySettingState string

const (
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_Compliant     ManagedDeviceEncryptionStateEncryptionPolicySettingState = "compliant"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_Conflict      ManagedDeviceEncryptionStateEncryptionPolicySettingState = "conflict"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_Error         ManagedDeviceEncryptionStateEncryptionPolicySettingState = "error"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_NonCompliant  ManagedDeviceEncryptionStateEncryptionPolicySettingState = "nonCompliant"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_NotApplicable ManagedDeviceEncryptionStateEncryptionPolicySettingState = "notApplicable"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_NotAssigned   ManagedDeviceEncryptionStateEncryptionPolicySettingState = "notAssigned"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_Remediated    ManagedDeviceEncryptionStateEncryptionPolicySettingState = "remediated"
	ManagedDeviceEncryptionStateEncryptionPolicySettingState_Unknown       ManagedDeviceEncryptionStateEncryptionPolicySettingState = "unknown"
)

func PossibleValuesForManagedDeviceEncryptionStateEncryptionPolicySettingState() []string {
	return []string{
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_Compliant),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_Conflict),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_Error),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_NonCompliant),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_NotApplicable),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_NotAssigned),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_Remediated),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingState_Unknown),
	}
}

func (s *ManagedDeviceEncryptionStateEncryptionPolicySettingState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceEncryptionStateEncryptionPolicySettingState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceEncryptionStateEncryptionPolicySettingState(input string) (*ManagedDeviceEncryptionStateEncryptionPolicySettingState, error) {
	vals := map[string]ManagedDeviceEncryptionStateEncryptionPolicySettingState{
		"compliant":     ManagedDeviceEncryptionStateEncryptionPolicySettingState_Compliant,
		"conflict":      ManagedDeviceEncryptionStateEncryptionPolicySettingState_Conflict,
		"error":         ManagedDeviceEncryptionStateEncryptionPolicySettingState_Error,
		"noncompliant":  ManagedDeviceEncryptionStateEncryptionPolicySettingState_NonCompliant,
		"notapplicable": ManagedDeviceEncryptionStateEncryptionPolicySettingState_NotApplicable,
		"notassigned":   ManagedDeviceEncryptionStateEncryptionPolicySettingState_NotAssigned,
		"remediated":    ManagedDeviceEncryptionStateEncryptionPolicySettingState_Remediated,
		"unknown":       ManagedDeviceEncryptionStateEncryptionPolicySettingState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateEncryptionPolicySettingState(input)
	return &out, nil
}
