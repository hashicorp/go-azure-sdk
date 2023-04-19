package codeversion

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssetProvisioningState string

const (
	AssetProvisioningStateCanceled  AssetProvisioningState = "Canceled"
	AssetProvisioningStateCreating  AssetProvisioningState = "Creating"
	AssetProvisioningStateDeleting  AssetProvisioningState = "Deleting"
	AssetProvisioningStateFailed    AssetProvisioningState = "Failed"
	AssetProvisioningStateSucceeded AssetProvisioningState = "Succeeded"
	AssetProvisioningStateUpdating  AssetProvisioningState = "Updating"
)

func PossibleValuesForAssetProvisioningState() []string {
	return []string{
		string(AssetProvisioningStateCanceled),
		string(AssetProvisioningStateCreating),
		string(AssetProvisioningStateDeleting),
		string(AssetProvisioningStateFailed),
		string(AssetProvisioningStateSucceeded),
		string(AssetProvisioningStateUpdating),
	}
}

func parseAssetProvisioningState(input string) (*AssetProvisioningState, error) {
	vals := map[string]AssetProvisioningState{
		"canceled":  AssetProvisioningStateCanceled,
		"creating":  AssetProvisioningStateCreating,
		"deleting":  AssetProvisioningStateDeleting,
		"failed":    AssetProvisioningStateFailed,
		"succeeded": AssetProvisioningStateSucceeded,
		"updating":  AssetProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssetProvisioningState(input)
	return &out, nil
}

type PendingUploadCredentialType string

const (
	PendingUploadCredentialTypeSAS PendingUploadCredentialType = "SAS"
)

func PossibleValuesForPendingUploadCredentialType() []string {
	return []string{
		string(PendingUploadCredentialTypeSAS),
	}
}

func parsePendingUploadCredentialType(input string) (*PendingUploadCredentialType, error) {
	vals := map[string]PendingUploadCredentialType{
		"sas": PendingUploadCredentialTypeSAS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PendingUploadCredentialType(input)
	return &out, nil
}

type PendingUploadType string

const (
	PendingUploadTypeNone                   PendingUploadType = "None"
	PendingUploadTypeTemporaryBlobReference PendingUploadType = "TemporaryBlobReference"
)

func PossibleValuesForPendingUploadType() []string {
	return []string{
		string(PendingUploadTypeNone),
		string(PendingUploadTypeTemporaryBlobReference),
	}
}

func parsePendingUploadType(input string) (*PendingUploadType, error) {
	vals := map[string]PendingUploadType{
		"none":                   PendingUploadTypeNone,
		"temporaryblobreference": PendingUploadTypeTemporaryBlobReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PendingUploadType(input)
	return &out, nil
}
