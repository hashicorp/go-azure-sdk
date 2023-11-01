package backupresourceencryptionconfigs

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionAtRestType string

const (
	EncryptionAtRestTypeCustomerManaged  EncryptionAtRestType = "CustomerManaged"
	EncryptionAtRestTypeInvalid          EncryptionAtRestType = "Invalid"
	EncryptionAtRestTypeMicrosoftManaged EncryptionAtRestType = "MicrosoftManaged"
)

func PossibleValuesForEncryptionAtRestType() []string {
	return []string{
		string(EncryptionAtRestTypeCustomerManaged),
		string(EncryptionAtRestTypeInvalid),
		string(EncryptionAtRestTypeMicrosoftManaged),
	}
}

func parseEncryptionAtRestType(input string) (*EncryptionAtRestType, error) {
	vals := map[string]EncryptionAtRestType{
		"customermanaged":  EncryptionAtRestTypeCustomerManaged,
		"invalid":          EncryptionAtRestTypeInvalid,
		"microsoftmanaged": EncryptionAtRestTypeMicrosoftManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionAtRestType(input)
	return &out, nil
}

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = "Disabled"
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = "Enabled"
	InfrastructureEncryptionStateInvalid  InfrastructureEncryptionState = "Invalid"
)

func PossibleValuesForInfrastructureEncryptionState() []string {
	return []string{
		string(InfrastructureEncryptionStateDisabled),
		string(InfrastructureEncryptionStateEnabled),
		string(InfrastructureEncryptionStateInvalid),
	}
}

func parseInfrastructureEncryptionState(input string) (*InfrastructureEncryptionState, error) {
	vals := map[string]InfrastructureEncryptionState{
		"disabled": InfrastructureEncryptionStateDisabled,
		"enabled":  InfrastructureEncryptionStateEnabled,
		"invalid":  InfrastructureEncryptionStateInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InfrastructureEncryptionState(input)
	return &out, nil
}

type LastUpdateStatus string

const (
	LastUpdateStatusFailed              LastUpdateStatus = "Failed"
	LastUpdateStatusFirstInitialization LastUpdateStatus = "FirstInitialization"
	LastUpdateStatusInitialized         LastUpdateStatus = "Initialized"
	LastUpdateStatusInvalid             LastUpdateStatus = "Invalid"
	LastUpdateStatusNotEnabled          LastUpdateStatus = "NotEnabled"
	LastUpdateStatusPartiallyFailed     LastUpdateStatus = "PartiallyFailed"
	LastUpdateStatusPartiallySucceeded  LastUpdateStatus = "PartiallySucceeded"
	LastUpdateStatusSucceeded           LastUpdateStatus = "Succeeded"
)

func PossibleValuesForLastUpdateStatus() []string {
	return []string{
		string(LastUpdateStatusFailed),
		string(LastUpdateStatusFirstInitialization),
		string(LastUpdateStatusInitialized),
		string(LastUpdateStatusInvalid),
		string(LastUpdateStatusNotEnabled),
		string(LastUpdateStatusPartiallyFailed),
		string(LastUpdateStatusPartiallySucceeded),
		string(LastUpdateStatusSucceeded),
	}
}

func parseLastUpdateStatus(input string) (*LastUpdateStatus, error) {
	vals := map[string]LastUpdateStatus{
		"failed":              LastUpdateStatusFailed,
		"firstinitialization": LastUpdateStatusFirstInitialization,
		"initialized":         LastUpdateStatusInitialized,
		"invalid":             LastUpdateStatusInvalid,
		"notenabled":          LastUpdateStatusNotEnabled,
		"partiallyfailed":     LastUpdateStatusPartiallyFailed,
		"partiallysucceeded":  LastUpdateStatusPartiallySucceeded,
		"succeeded":           LastUpdateStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LastUpdateStatus(input)
	return &out, nil
}
