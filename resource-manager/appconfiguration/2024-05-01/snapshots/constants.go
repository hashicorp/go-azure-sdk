package snapshots

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompositionType string

const (
	CompositionTypeKey      CompositionType = "Key"
	CompositionTypeKeyLabel CompositionType = "Key_Label"
)

func PossibleValuesForCompositionType() []string {
	return []string{
		string(CompositionTypeKey),
		string(CompositionTypeKeyLabel),
	}
}

func (s *CompositionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompositionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompositionType(input string) (*CompositionType, error) {
	vals := map[string]CompositionType{
		"key":       CompositionTypeKey,
		"key_label": CompositionTypeKeyLabel,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompositionType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
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
		"canceled":  ProvisioningStateCanceled,
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
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

type SnapshotStatus string

const (
	SnapshotStatusArchived     SnapshotStatus = "Archived"
	SnapshotStatusFailed       SnapshotStatus = "Failed"
	SnapshotStatusProvisioning SnapshotStatus = "Provisioning"
	SnapshotStatusReady        SnapshotStatus = "Ready"
)

func PossibleValuesForSnapshotStatus() []string {
	return []string{
		string(SnapshotStatusArchived),
		string(SnapshotStatusFailed),
		string(SnapshotStatusProvisioning),
		string(SnapshotStatusReady),
	}
}

func (s *SnapshotStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSnapshotStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSnapshotStatus(input string) (*SnapshotStatus, error) {
	vals := map[string]SnapshotStatus{
		"archived":     SnapshotStatusArchived,
		"failed":       SnapshotStatusFailed,
		"provisioning": SnapshotStatusProvisioning,
		"ready":        SnapshotStatusReady,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SnapshotStatus(input)
	return &out, nil
}
