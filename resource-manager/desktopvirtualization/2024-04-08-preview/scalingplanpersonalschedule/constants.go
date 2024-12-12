package scalingplanpersonalschedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
	}
}

func (s *DayOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDayOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDayOfWeek(input string) (*DayOfWeek, error) {
	vals := map[string]DayOfWeek{
		"friday":    DayOfWeekFriday,
		"monday":    DayOfWeekMonday,
		"saturday":  DayOfWeekSaturday,
		"sunday":    DayOfWeekSunday,
		"thursday":  DayOfWeekThursday,
		"tuesday":   DayOfWeekTuesday,
		"wednesday": DayOfWeekWednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DayOfWeek(input)
	return &out, nil
}

type SessionHandlingOperation string

const (
	SessionHandlingOperationDeallocate SessionHandlingOperation = "Deallocate"
	SessionHandlingOperationHibernate  SessionHandlingOperation = "Hibernate"
	SessionHandlingOperationNone       SessionHandlingOperation = "None"
)

func PossibleValuesForSessionHandlingOperation() []string {
	return []string{
		string(SessionHandlingOperationDeallocate),
		string(SessionHandlingOperationHibernate),
		string(SessionHandlingOperationNone),
	}
}

func (s *SessionHandlingOperation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSessionHandlingOperation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSessionHandlingOperation(input string) (*SessionHandlingOperation, error) {
	vals := map[string]SessionHandlingOperation{
		"deallocate": SessionHandlingOperationDeallocate,
		"hibernate":  SessionHandlingOperationHibernate,
		"none":       SessionHandlingOperationNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SessionHandlingOperation(input)
	return &out, nil
}

type SetStartVMOnConnect string

const (
	SetStartVMOnConnectDisable SetStartVMOnConnect = "Disable"
	SetStartVMOnConnectEnable  SetStartVMOnConnect = "Enable"
)

func PossibleValuesForSetStartVMOnConnect() []string {
	return []string{
		string(SetStartVMOnConnectDisable),
		string(SetStartVMOnConnectEnable),
	}
}

func (s *SetStartVMOnConnect) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSetStartVMOnConnect(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSetStartVMOnConnect(input string) (*SetStartVMOnConnect, error) {
	vals := map[string]SetStartVMOnConnect{
		"disable": SetStartVMOnConnectDisable,
		"enable":  SetStartVMOnConnectEnable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SetStartVMOnConnect(input)
	return &out, nil
}

type StartupBehavior string

const (
	StartupBehaviorAll              StartupBehavior = "All"
	StartupBehaviorNone             StartupBehavior = "None"
	StartupBehaviorWithAssignedUser StartupBehavior = "WithAssignedUser"
)

func PossibleValuesForStartupBehavior() []string {
	return []string{
		string(StartupBehaviorAll),
		string(StartupBehaviorNone),
		string(StartupBehaviorWithAssignedUser),
	}
}

func (s *StartupBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStartupBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStartupBehavior(input string) (*StartupBehavior, error) {
	vals := map[string]StartupBehavior{
		"all":              StartupBehaviorAll,
		"none":             StartupBehaviorNone,
		"withassigneduser": StartupBehaviorWithAssignedUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StartupBehavior(input)
	return &out, nil
}
