package galleryimageversions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregatedReplicationState string

const (
	AggregatedReplicationStateCompleted  AggregatedReplicationState = "Completed"
	AggregatedReplicationStateFailed     AggregatedReplicationState = "Failed"
	AggregatedReplicationStateInProgress AggregatedReplicationState = "InProgress"
	AggregatedReplicationStateUnknown    AggregatedReplicationState = "Unknown"
)

func PossibleValuesForAggregatedReplicationState() []string {
	return []string{
		string(AggregatedReplicationStateCompleted),
		string(AggregatedReplicationStateFailed),
		string(AggregatedReplicationStateInProgress),
		string(AggregatedReplicationStateUnknown),
	}
}

func (s *AggregatedReplicationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAggregatedReplicationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAggregatedReplicationState(input string) (*AggregatedReplicationState, error) {
	vals := map[string]AggregatedReplicationState{
		"completed":  AggregatedReplicationStateCompleted,
		"failed":     AggregatedReplicationStateFailed,
		"inprogress": AggregatedReplicationStateInProgress,
		"unknown":    AggregatedReplicationStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AggregatedReplicationState(input)
	return &out, nil
}

type HostCaching string

const (
	HostCachingNone      HostCaching = "None"
	HostCachingReadOnly  HostCaching = "ReadOnly"
	HostCachingReadWrite HostCaching = "ReadWrite"
)

func PossibleValuesForHostCaching() []string {
	return []string{
		string(HostCachingNone),
		string(HostCachingReadOnly),
		string(HostCachingReadWrite),
	}
}

func (s *HostCaching) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHostCaching(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHostCaching(input string) (*HostCaching, error) {
	vals := map[string]HostCaching{
		"none":      HostCachingNone,
		"readonly":  HostCachingReadOnly,
		"readwrite": HostCachingReadWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HostCaching(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMigrating ProvisioningState = "Migrating"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMigrating),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"migrating": ProvisioningStateMigrating,
		"succeeded": ProvisioningStateSucceeded,
		"updating":  ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type ReplicationMode string

const (
	ReplicationModeFull    ReplicationMode = "Full"
	ReplicationModeShallow ReplicationMode = "Shallow"
)

func PossibleValuesForReplicationMode() []string {
	return []string{
		string(ReplicationModeFull),
		string(ReplicationModeShallow),
	}
}

func (s *ReplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicationMode(input string) (*ReplicationMode, error) {
	vals := map[string]ReplicationMode{
		"full":    ReplicationModeFull,
		"shallow": ReplicationModeShallow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationMode(input)
	return &out, nil
}

type ReplicationState string

const (
	ReplicationStateCompleted   ReplicationState = "Completed"
	ReplicationStateFailed      ReplicationState = "Failed"
	ReplicationStateReplicating ReplicationState = "Replicating"
	ReplicationStateUnknown     ReplicationState = "Unknown"
)

func PossibleValuesForReplicationState() []string {
	return []string{
		string(ReplicationStateCompleted),
		string(ReplicationStateFailed),
		string(ReplicationStateReplicating),
		string(ReplicationStateUnknown),
	}
}

func (s *ReplicationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicationState(input string) (*ReplicationState, error) {
	vals := map[string]ReplicationState{
		"completed":   ReplicationStateCompleted,
		"failed":      ReplicationStateFailed,
		"replicating": ReplicationStateReplicating,
		"unknown":     ReplicationStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationState(input)
	return &out, nil
}

type ReplicationStatusTypes string

const (
	ReplicationStatusTypesReplicationStatus ReplicationStatusTypes = "ReplicationStatus"
)

func PossibleValuesForReplicationStatusTypes() []string {
	return []string{
		string(ReplicationStatusTypesReplicationStatus),
	}
}

func (s *ReplicationStatusTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReplicationStatusTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReplicationStatusTypes(input string) (*ReplicationStatusTypes, error) {
	vals := map[string]ReplicationStatusTypes{
		"replicationstatus": ReplicationStatusTypesReplicationStatus,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicationStatusTypes(input)
	return &out, nil
}

type StorageAccountType string

const (
	StorageAccountTypePremiumLRS  StorageAccountType = "Premium_LRS"
	StorageAccountTypeStandardLRS StorageAccountType = "Standard_LRS"
	StorageAccountTypeStandardZRS StorageAccountType = "Standard_ZRS"
)

func PossibleValuesForStorageAccountType() []string {
	return []string{
		string(StorageAccountTypePremiumLRS),
		string(StorageAccountTypeStandardLRS),
		string(StorageAccountTypeStandardZRS),
	}
}

func (s *StorageAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStorageAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStorageAccountType(input string) (*StorageAccountType, error) {
	vals := map[string]StorageAccountType{
		"premium_lrs":  StorageAccountTypePremiumLRS,
		"standard_lrs": StorageAccountTypeStandardLRS,
		"standard_zrs": StorageAccountTypeStandardZRS,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StorageAccountType(input)
	return &out, nil
}
