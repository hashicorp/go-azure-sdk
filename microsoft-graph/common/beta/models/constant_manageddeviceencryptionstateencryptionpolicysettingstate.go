package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateEncryptionPolicySettingState string

const (
	ManagedDeviceEncryptionStateEncryptionPolicySettingStatecompliant     ManagedDeviceEncryptionStateEncryptionPolicySettingState = "Compliant"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStateconflict      ManagedDeviceEncryptionStateEncryptionPolicySettingState = "Conflict"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStateerror         ManagedDeviceEncryptionStateEncryptionPolicySettingState = "Error"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStatenonCompliant  ManagedDeviceEncryptionStateEncryptionPolicySettingState = "NonCompliant"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStatenotApplicable ManagedDeviceEncryptionStateEncryptionPolicySettingState = "NotApplicable"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStatenotAssigned   ManagedDeviceEncryptionStateEncryptionPolicySettingState = "NotAssigned"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStateremediated    ManagedDeviceEncryptionStateEncryptionPolicySettingState = "Remediated"
	ManagedDeviceEncryptionStateEncryptionPolicySettingStateunknown       ManagedDeviceEncryptionStateEncryptionPolicySettingState = "Unknown"
)

func PossibleValuesForManagedDeviceEncryptionStateEncryptionPolicySettingState() []string {
	return []string{
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStatecompliant),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStateconflict),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStateerror),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStatenonCompliant),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStatenotApplicable),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStatenotAssigned),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStateremediated),
		string(ManagedDeviceEncryptionStateEncryptionPolicySettingStateunknown),
	}
}

func parseManagedDeviceEncryptionStateEncryptionPolicySettingState(input string) (*ManagedDeviceEncryptionStateEncryptionPolicySettingState, error) {
	vals := map[string]ManagedDeviceEncryptionStateEncryptionPolicySettingState{
		"compliant":     ManagedDeviceEncryptionStateEncryptionPolicySettingStatecompliant,
		"conflict":      ManagedDeviceEncryptionStateEncryptionPolicySettingStateconflict,
		"error":         ManagedDeviceEncryptionStateEncryptionPolicySettingStateerror,
		"noncompliant":  ManagedDeviceEncryptionStateEncryptionPolicySettingStatenonCompliant,
		"notapplicable": ManagedDeviceEncryptionStateEncryptionPolicySettingStatenotApplicable,
		"notassigned":   ManagedDeviceEncryptionStateEncryptionPolicySettingStatenotAssigned,
		"remediated":    ManagedDeviceEncryptionStateEncryptionPolicySettingStateremediated,
		"unknown":       ManagedDeviceEncryptionStateEncryptionPolicySettingStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateEncryptionPolicySettingState(input)
	return &out, nil
}
