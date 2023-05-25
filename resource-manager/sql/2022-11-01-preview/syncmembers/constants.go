package syncmembers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncDirection string

const (
	SyncDirectionBidirectional     SyncDirection = "Bidirectional"
	SyncDirectionOneWayHubToMember SyncDirection = "OneWayHubToMember"
	SyncDirectionOneWayMemberToHub SyncDirection = "OneWayMemberToHub"
)

func PossibleValuesForSyncDirection() []string {
	return []string{
		string(SyncDirectionBidirectional),
		string(SyncDirectionOneWayHubToMember),
		string(SyncDirectionOneWayMemberToHub),
	}
}

func (s *SyncDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncDirection(input string) (*SyncDirection, error) {
	vals := map[string]SyncDirection{
		"bidirectional":     SyncDirectionBidirectional,
		"onewayhubtomember": SyncDirectionOneWayHubToMember,
		"onewaymembertohub": SyncDirectionOneWayMemberToHub,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncDirection(input)
	return &out, nil
}

type SyncMemberDbType string

const (
	SyncMemberDbTypeAzureSqlDatabase  SyncMemberDbType = "AzureSqlDatabase"
	SyncMemberDbTypeSqlServerDatabase SyncMemberDbType = "SqlServerDatabase"
)

func PossibleValuesForSyncMemberDbType() []string {
	return []string{
		string(SyncMemberDbTypeAzureSqlDatabase),
		string(SyncMemberDbTypeSqlServerDatabase),
	}
}

func (s *SyncMemberDbType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncMemberDbType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncMemberDbType(input string) (*SyncMemberDbType, error) {
	vals := map[string]SyncMemberDbType{
		"azuresqldatabase":  SyncMemberDbTypeAzureSqlDatabase,
		"sqlserverdatabase": SyncMemberDbTypeSqlServerDatabase,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncMemberDbType(input)
	return &out, nil
}

type SyncMemberState string

const (
	SyncMemberStateDeProvisionFailed         SyncMemberState = "DeProvisionFailed"
	SyncMemberStateDeProvisioned             SyncMemberState = "DeProvisioned"
	SyncMemberStateDeProvisioning            SyncMemberState = "DeProvisioning"
	SyncMemberStateDisabledBackupRestore     SyncMemberState = "DisabledBackupRestore"
	SyncMemberStateDisabledTombstoneCleanup  SyncMemberState = "DisabledTombstoneCleanup"
	SyncMemberStateProvisionFailed           SyncMemberState = "ProvisionFailed"
	SyncMemberStateProvisioned               SyncMemberState = "Provisioned"
	SyncMemberStateProvisioning              SyncMemberState = "Provisioning"
	SyncMemberStateReprovisionFailed         SyncMemberState = "ReprovisionFailed"
	SyncMemberStateReprovisioning            SyncMemberState = "Reprovisioning"
	SyncMemberStateSyncCancelled             SyncMemberState = "SyncCancelled"
	SyncMemberStateSyncCancelling            SyncMemberState = "SyncCancelling"
	SyncMemberStateSyncFailed                SyncMemberState = "SyncFailed"
	SyncMemberStateSyncInProgress            SyncMemberState = "SyncInProgress"
	SyncMemberStateSyncSucceeded             SyncMemberState = "SyncSucceeded"
	SyncMemberStateSyncSucceededWithWarnings SyncMemberState = "SyncSucceededWithWarnings"
	SyncMemberStateUnProvisioned             SyncMemberState = "UnProvisioned"
	SyncMemberStateUnReprovisioned           SyncMemberState = "UnReprovisioned"
)

func PossibleValuesForSyncMemberState() []string {
	return []string{
		string(SyncMemberStateDeProvisionFailed),
		string(SyncMemberStateDeProvisioned),
		string(SyncMemberStateDeProvisioning),
		string(SyncMemberStateDisabledBackupRestore),
		string(SyncMemberStateDisabledTombstoneCleanup),
		string(SyncMemberStateProvisionFailed),
		string(SyncMemberStateProvisioned),
		string(SyncMemberStateProvisioning),
		string(SyncMemberStateReprovisionFailed),
		string(SyncMemberStateReprovisioning),
		string(SyncMemberStateSyncCancelled),
		string(SyncMemberStateSyncCancelling),
		string(SyncMemberStateSyncFailed),
		string(SyncMemberStateSyncInProgress),
		string(SyncMemberStateSyncSucceeded),
		string(SyncMemberStateSyncSucceededWithWarnings),
		string(SyncMemberStateUnProvisioned),
		string(SyncMemberStateUnReprovisioned),
	}
}

func (s *SyncMemberState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncMemberState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncMemberState(input string) (*SyncMemberState, error) {
	vals := map[string]SyncMemberState{
		"deprovisionfailed":         SyncMemberStateDeProvisionFailed,
		"deprovisioned":             SyncMemberStateDeProvisioned,
		"deprovisioning":            SyncMemberStateDeProvisioning,
		"disabledbackuprestore":     SyncMemberStateDisabledBackupRestore,
		"disabledtombstonecleanup":  SyncMemberStateDisabledTombstoneCleanup,
		"provisionfailed":           SyncMemberStateProvisionFailed,
		"provisioned":               SyncMemberStateProvisioned,
		"provisioning":              SyncMemberStateProvisioning,
		"reprovisionfailed":         SyncMemberStateReprovisionFailed,
		"reprovisioning":            SyncMemberStateReprovisioning,
		"synccancelled":             SyncMemberStateSyncCancelled,
		"synccancelling":            SyncMemberStateSyncCancelling,
		"syncfailed":                SyncMemberStateSyncFailed,
		"syncinprogress":            SyncMemberStateSyncInProgress,
		"syncsucceeded":             SyncMemberStateSyncSucceeded,
		"syncsucceededwithwarnings": SyncMemberStateSyncSucceededWithWarnings,
		"unprovisioned":             SyncMemberStateUnProvisioned,
		"unreprovisioned":           SyncMemberStateUnReprovisioned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncMemberState(input)
	return &out, nil
}
