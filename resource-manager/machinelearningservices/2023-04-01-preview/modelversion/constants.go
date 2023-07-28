package modelversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *AssetProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssetProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type AutoDeleteCondition string

const (
	AutoDeleteConditionCreatedGreaterThan      AutoDeleteCondition = "CreatedGreaterThan"
	AutoDeleteConditionLastAccessedGreaterThan AutoDeleteCondition = "LastAccessedGreaterThan"
)

func PossibleValuesForAutoDeleteCondition() []string {
	return []string{
		string(AutoDeleteConditionCreatedGreaterThan),
		string(AutoDeleteConditionLastAccessedGreaterThan),
	}
}

func (s *AutoDeleteCondition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutoDeleteCondition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutoDeleteCondition(input string) (*AutoDeleteCondition, error) {
	vals := map[string]AutoDeleteCondition{
		"createdgreaterthan":      AutoDeleteConditionCreatedGreaterThan,
		"lastaccessedgreaterthan": AutoDeleteConditionLastAccessedGreaterThan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutoDeleteCondition(input)
	return &out, nil
}

type BaseEnvironmentSourceType string

const (
	BaseEnvironmentSourceTypeEnvironmentAsset BaseEnvironmentSourceType = "EnvironmentAsset"
)

func PossibleValuesForBaseEnvironmentSourceType() []string {
	return []string{
		string(BaseEnvironmentSourceTypeEnvironmentAsset),
	}
}

func (s *BaseEnvironmentSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBaseEnvironmentSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBaseEnvironmentSourceType(input string) (*BaseEnvironmentSourceType, error) {
	vals := map[string]BaseEnvironmentSourceType{
		"environmentasset": BaseEnvironmentSourceTypeEnvironmentAsset,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BaseEnvironmentSourceType(input)
	return &out, nil
}

type InferencingServerType string

const (
	InferencingServerTypeAzureMLBatch  InferencingServerType = "AzureMLBatch"
	InferencingServerTypeAzureMLOnline InferencingServerType = "AzureMLOnline"
	InferencingServerTypeCustom        InferencingServerType = "Custom"
	InferencingServerTypeTriton        InferencingServerType = "Triton"
)

func PossibleValuesForInferencingServerType() []string {
	return []string{
		string(InferencingServerTypeAzureMLBatch),
		string(InferencingServerTypeAzureMLOnline),
		string(InferencingServerTypeCustom),
		string(InferencingServerTypeTriton),
	}
}

func (s *InferencingServerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInferencingServerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInferencingServerType(input string) (*InferencingServerType, error) {
	vals := map[string]InferencingServerType{
		"azuremlbatch":  InferencingServerTypeAzureMLBatch,
		"azuremlonline": InferencingServerTypeAzureMLOnline,
		"custom":        InferencingServerTypeCustom,
		"triton":        InferencingServerTypeTriton,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InferencingServerType(input)
	return &out, nil
}

type InputPathType string

const (
	InputPathTypePathId      InputPathType = "PathId"
	InputPathTypePathVersion InputPathType = "PathVersion"
	InputPathTypeUrl         InputPathType = "Url"
)

func PossibleValuesForInputPathType() []string {
	return []string{
		string(InputPathTypePathId),
		string(InputPathTypePathVersion),
		string(InputPathTypeUrl),
	}
}

func (s *InputPathType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInputPathType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInputPathType(input string) (*InputPathType, error) {
	vals := map[string]InputPathType{
		"pathid":      InputPathTypePathId,
		"pathversion": InputPathTypePathVersion,
		"url":         InputPathTypeUrl,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InputPathType(input)
	return &out, nil
}

type ListViewType string

const (
	ListViewTypeActiveOnly   ListViewType = "ActiveOnly"
	ListViewTypeAll          ListViewType = "All"
	ListViewTypeArchivedOnly ListViewType = "ArchivedOnly"
)

func PossibleValuesForListViewType() []string {
	return []string{
		string(ListViewTypeActiveOnly),
		string(ListViewTypeAll),
		string(ListViewTypeArchivedOnly),
	}
}

func (s *ListViewType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseListViewType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseListViewType(input string) (*ListViewType, error) {
	vals := map[string]ListViewType{
		"activeonly":   ListViewTypeActiveOnly,
		"all":          ListViewTypeAll,
		"archivedonly": ListViewTypeArchivedOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ListViewType(input)
	return &out, nil
}

type PackageBuildState string

const (
	PackageBuildStateFailed     PackageBuildState = "Failed"
	PackageBuildStateNotStarted PackageBuildState = "NotStarted"
	PackageBuildStateRunning    PackageBuildState = "Running"
	PackageBuildStateSucceeded  PackageBuildState = "Succeeded"
)

func PossibleValuesForPackageBuildState() []string {
	return []string{
		string(PackageBuildStateFailed),
		string(PackageBuildStateNotStarted),
		string(PackageBuildStateRunning),
		string(PackageBuildStateSucceeded),
	}
}

func (s *PackageBuildState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePackageBuildState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePackageBuildState(input string) (*PackageBuildState, error) {
	vals := map[string]PackageBuildState{
		"failed":     PackageBuildStateFailed,
		"notstarted": PackageBuildStateNotStarted,
		"running":    PackageBuildStateRunning,
		"succeeded":  PackageBuildStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PackageBuildState(input)
	return &out, nil
}

type PackageInputDeliveryMode string

const (
	PackageInputDeliveryModeDownload      PackageInputDeliveryMode = "Download"
	PackageInputDeliveryModeReadOnlyMount PackageInputDeliveryMode = "ReadOnlyMount"
)

func PossibleValuesForPackageInputDeliveryMode() []string {
	return []string{
		string(PackageInputDeliveryModeDownload),
		string(PackageInputDeliveryModeReadOnlyMount),
	}
}

func (s *PackageInputDeliveryMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePackageInputDeliveryMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePackageInputDeliveryMode(input string) (*PackageInputDeliveryMode, error) {
	vals := map[string]PackageInputDeliveryMode{
		"download":      PackageInputDeliveryModeDownload,
		"readonlymount": PackageInputDeliveryModeReadOnlyMount,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PackageInputDeliveryMode(input)
	return &out, nil
}

type PackageInputType string

const (
	PackageInputTypeUriFile   PackageInputType = "UriFile"
	PackageInputTypeUriFolder PackageInputType = "UriFolder"
)

func PossibleValuesForPackageInputType() []string {
	return []string{
		string(PackageInputTypeUriFile),
		string(PackageInputTypeUriFolder),
	}
}

func (s *PackageInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePackageInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePackageInputType(input string) (*PackageInputType, error) {
	vals := map[string]PackageInputType{
		"urifile":   PackageInputTypeUriFile,
		"urifolder": PackageInputTypeUriFolder,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PackageInputType(input)
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

func (s *PendingUploadCredentialType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePendingUploadCredentialType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

func (s *PendingUploadType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePendingUploadType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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

type ProtectionLevel string

const (
	ProtectionLevelAll  ProtectionLevel = "All"
	ProtectionLevelNone ProtectionLevel = "None"
)

func PossibleValuesForProtectionLevel() []string {
	return []string{
		string(ProtectionLevelAll),
		string(ProtectionLevelNone),
	}
}

func (s *ProtectionLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectionLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectionLevel(input string) (*ProtectionLevel, error) {
	vals := map[string]ProtectionLevel{
		"all":  ProtectionLevelAll,
		"none": ProtectionLevelNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectionLevel(input)
	return &out, nil
}
