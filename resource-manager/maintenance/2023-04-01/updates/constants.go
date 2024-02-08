package updates

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImpactType string

const (
	ImpactTypeFreeze   ImpactType = "Freeze"
	ImpactTypeNone     ImpactType = "None"
	ImpactTypeRedeploy ImpactType = "Redeploy"
	ImpactTypeRestart  ImpactType = "Restart"
)

func PossibleValuesForImpactType() []string {
	return []string{
		string(ImpactTypeFreeze),
		string(ImpactTypeNone),
		string(ImpactTypeRedeploy),
		string(ImpactTypeRestart),
	}
}

func (s *ImpactType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImpactType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImpactType(input string) (*ImpactType, error) {
	vals := map[string]ImpactType{
		"freeze":   ImpactTypeFreeze,
		"none":     ImpactTypeNone,
		"redeploy": ImpactTypeRedeploy,
		"restart":  ImpactTypeRestart,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImpactType(input)
	return &out, nil
}

type MaintenanceScope string

const (
	MaintenanceScopeExtension          MaintenanceScope = "Extension"
	MaintenanceScopeHost               MaintenanceScope = "Host"
	MaintenanceScopeInGuestPatch       MaintenanceScope = "InGuestPatch"
	MaintenanceScopeOSImage            MaintenanceScope = "OSImage"
	MaintenanceScopeResource           MaintenanceScope = "Resource"
	MaintenanceScopeSQLDB              MaintenanceScope = "SQLDB"
	MaintenanceScopeSQLManagedInstance MaintenanceScope = "SQLManagedInstance"
)

func PossibleValuesForMaintenanceScope() []string {
	return []string{
		string(MaintenanceScopeExtension),
		string(MaintenanceScopeHost),
		string(MaintenanceScopeInGuestPatch),
		string(MaintenanceScopeOSImage),
		string(MaintenanceScopeResource),
		string(MaintenanceScopeSQLDB),
		string(MaintenanceScopeSQLManagedInstance),
	}
}

func (s *MaintenanceScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMaintenanceScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMaintenanceScope(input string) (*MaintenanceScope, error) {
	vals := map[string]MaintenanceScope{
		"extension":          MaintenanceScopeExtension,
		"host":               MaintenanceScopeHost,
		"inguestpatch":       MaintenanceScopeInGuestPatch,
		"osimage":            MaintenanceScopeOSImage,
		"resource":           MaintenanceScopeResource,
		"sqldb":              MaintenanceScopeSQLDB,
		"sqlmanagedinstance": MaintenanceScopeSQLManagedInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MaintenanceScope(input)
	return &out, nil
}

type UpdateStatus string

const (
	UpdateStatusCompleted  UpdateStatus = "Completed"
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusPending    UpdateStatus = "Pending"
	UpdateStatusRetryLater UpdateStatus = "RetryLater"
	UpdateStatusRetryNow   UpdateStatus = "RetryNow"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusCompleted),
		string(UpdateStatusInProgress),
		string(UpdateStatusPending),
		string(UpdateStatusRetryLater),
		string(UpdateStatusRetryNow),
	}
}

func (s *UpdateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateStatus(input string) (*UpdateStatus, error) {
	vals := map[string]UpdateStatus{
		"completed":  UpdateStatusCompleted,
		"inprogress": UpdateStatusInProgress,
		"pending":    UpdateStatusPending,
		"retrylater": UpdateStatusRetryLater,
		"retrynow":   UpdateStatusRetryNow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateStatus(input)
	return &out, nil
}
