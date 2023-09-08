package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionStateEncryptionReadinessState string

const (
	ManagedDeviceEncryptionStateEncryptionReadinessStatenotReady ManagedDeviceEncryptionStateEncryptionReadinessState = "NotReady"
	ManagedDeviceEncryptionStateEncryptionReadinessStateready    ManagedDeviceEncryptionStateEncryptionReadinessState = "Ready"
)

func PossibleValuesForManagedDeviceEncryptionStateEncryptionReadinessState() []string {
	return []string{
		string(ManagedDeviceEncryptionStateEncryptionReadinessStatenotReady),
		string(ManagedDeviceEncryptionStateEncryptionReadinessStateready),
	}
}

func parseManagedDeviceEncryptionStateEncryptionReadinessState(input string) (*ManagedDeviceEncryptionStateEncryptionReadinessState, error) {
	vals := map[string]ManagedDeviceEncryptionStateEncryptionReadinessState{
		"notready": ManagedDeviceEncryptionStateEncryptionReadinessStatenotReady,
		"ready":    ManagedDeviceEncryptionStateEncryptionReadinessStateready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceEncryptionStateEncryptionReadinessState(input)
	return &out, nil
}
