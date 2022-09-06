package volumegroups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Action string

const (
	ActionAllow Action = "Allow"
)

func PossibleValuesForAction() []string {
	return []string{
		string(ActionAllow),
	}
}

func parseAction(input string) (*Action, error) {
	vals := map[string]Action{
		"allow": ActionAllow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Action(input)
	return &out, nil
}

type EncryptionType string

const (
	EncryptionTypeEncryptionAtRestWithPlatformKey EncryptionType = "EncryptionAtRestWithPlatformKey"
)

func PossibleValuesForEncryptionType() []string {
	return []string{
		string(EncryptionTypeEncryptionAtRestWithPlatformKey),
	}
}

func parseEncryptionType(input string) (*EncryptionType, error) {
	vals := map[string]EncryptionType{
		"encryptionatrestwithplatformkey": EncryptionTypeEncryptionAtRestWithPlatformKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EncryptionType(input)
	return &out, nil
}

type ProvisioningStates string

const (
	ProvisioningStatesCanceled  ProvisioningStates = "Canceled"
	ProvisioningStatesCreating  ProvisioningStates = "Creating"
	ProvisioningStatesDeleting  ProvisioningStates = "Deleting"
	ProvisioningStatesFailed    ProvisioningStates = "Failed"
	ProvisioningStatesInvalid   ProvisioningStates = "Invalid"
	ProvisioningStatesPending   ProvisioningStates = "Pending"
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	ProvisioningStatesUpdating  ProvisioningStates = "Updating"
)

func PossibleValuesForProvisioningStates() []string {
	return []string{
		string(ProvisioningStatesCanceled),
		string(ProvisioningStatesCreating),
		string(ProvisioningStatesDeleting),
		string(ProvisioningStatesFailed),
		string(ProvisioningStatesInvalid),
		string(ProvisioningStatesPending),
		string(ProvisioningStatesSucceeded),
		string(ProvisioningStatesUpdating),
	}
}

func parseProvisioningStates(input string) (*ProvisioningStates, error) {
	vals := map[string]ProvisioningStates{
		"canceled":  ProvisioningStatesCanceled,
		"creating":  ProvisioningStatesCreating,
		"deleting":  ProvisioningStatesDeleting,
		"failed":    ProvisioningStatesFailed,
		"invalid":   ProvisioningStatesInvalid,
		"pending":   ProvisioningStatesPending,
		"succeeded": ProvisioningStatesSucceeded,
		"updating":  ProvisioningStatesUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningStates(input)
	return &out, nil
}

type State string

const (
	StateDeprovisioning       State = "deprovisioning"
	StateFailed               State = "failed"
	StateNetworkSourceDeleted State = "networkSourceDeleted"
	StateProvisioning         State = "provisioning"
	StateSucceeded            State = "succeeded"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateDeprovisioning),
		string(StateFailed),
		string(StateNetworkSourceDeleted),
		string(StateProvisioning),
		string(StateSucceeded),
	}
}

func parseState(input string) (*State, error) {
	vals := map[string]State{
		"deprovisioning":       StateDeprovisioning,
		"failed":               StateFailed,
		"networksourcedeleted": StateNetworkSourceDeleted,
		"provisioning":         StateProvisioning,
		"succeeded":            StateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := State(input)
	return &out, nil
}

type StorageTargetType string

const (
	StorageTargetTypeIscsi StorageTargetType = "Iscsi"
	StorageTargetTypeNone  StorageTargetType = "None"
)

func PossibleValuesForStorageTargetType() []string {
	return []string{
		string(StorageTargetTypeIscsi),
		string(StorageTargetTypeNone),
	}
}

func parseStorageTargetType(input string) (*StorageTargetType, error) {
	vals := map[string]StorageTargetType{
		"iscsi": StorageTargetTypeIscsi,
		"none":  StorageTargetTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageTargetType(input)
	return &out, nil
}
