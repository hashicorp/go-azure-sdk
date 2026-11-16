package maintenanceevents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MaintenanceEventStatus string

const (
	MaintenanceEventStatusCanceled    MaintenanceEventStatus = "Canceled"
	MaintenanceEventStatusComplete    MaintenanceEventStatus = "Complete"
	MaintenanceEventStatusInProgress  MaintenanceEventStatus = "InProgress"
	MaintenanceEventStatusPlanned     MaintenanceEventStatus = "Planned"
	MaintenanceEventStatusRescheduled MaintenanceEventStatus = "Rescheduled"
)

func PossibleValuesForMaintenanceEventStatus() []string {
	return []string{
		string(MaintenanceEventStatusCanceled),
		string(MaintenanceEventStatusComplete),
		string(MaintenanceEventStatusInProgress),
		string(MaintenanceEventStatusPlanned),
		string(MaintenanceEventStatusRescheduled),
	}
}

func (s *MaintenanceEventStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMaintenanceEventStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMaintenanceEventStatus(input string) (*MaintenanceEventStatus, error) {
	vals := map[string]MaintenanceEventStatus{
		"canceled":    MaintenanceEventStatusCanceled,
		"complete":    MaintenanceEventStatusComplete,
		"inprogress":  MaintenanceEventStatusInProgress,
		"planned":     MaintenanceEventStatusPlanned,
		"rescheduled": MaintenanceEventStatusRescheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MaintenanceEventStatus(input)
	return &out, nil
}

type MaintenanceEventStatusFilter string

const (
	MaintenanceEventStatusFilterPast     MaintenanceEventStatusFilter = "Past"
	MaintenanceEventStatusFilterUpcoming MaintenanceEventStatusFilter = "Upcoming"
)

func PossibleValuesForMaintenanceEventStatusFilter() []string {
	return []string{
		string(MaintenanceEventStatusFilterPast),
		string(MaintenanceEventStatusFilterUpcoming),
	}
}

func (s *MaintenanceEventStatusFilter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMaintenanceEventStatusFilter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMaintenanceEventStatusFilter(input string) (*MaintenanceEventStatusFilter, error) {
	vals := map[string]MaintenanceEventStatusFilter{
		"past":     MaintenanceEventStatusFilterPast,
		"upcoming": MaintenanceEventStatusFilterUpcoming,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MaintenanceEventStatusFilter(input)
	return &out, nil
}

type MaintenanceType string

const (
	MaintenanceTypePlannedMaintenance MaintenanceType = "PlannedMaintenance"
)

func PossibleValuesForMaintenanceType() []string {
	return []string{
		string(MaintenanceTypePlannedMaintenance),
	}
}

func (s *MaintenanceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMaintenanceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMaintenanceType(input string) (*MaintenanceType, error) {
	vals := map[string]MaintenanceType{
		"plannedmaintenance": MaintenanceTypePlannedMaintenance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MaintenanceType(input)
	return &out, nil
}
