package syncgroups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncConflictResolutionPolicy string

const (
	SyncConflictResolutionPolicyHubWin    SyncConflictResolutionPolicy = "HubWin"
	SyncConflictResolutionPolicyMemberWin SyncConflictResolutionPolicy = "MemberWin"
)

func PossibleValuesForSyncConflictResolutionPolicy() []string {
	return []string{
		string(SyncConflictResolutionPolicyHubWin),
		string(SyncConflictResolutionPolicyMemberWin),
	}
}

func (s *SyncConflictResolutionPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncConflictResolutionPolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncConflictResolutionPolicy(input string) (*SyncConflictResolutionPolicy, error) {
	vals := map[string]SyncConflictResolutionPolicy{
		"hubwin":    SyncConflictResolutionPolicyHubWin,
		"memberwin": SyncConflictResolutionPolicyMemberWin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncConflictResolutionPolicy(input)
	return &out, nil
}

type SyncGroupLogType string

const (
	SyncGroupLogTypeAll     SyncGroupLogType = "All"
	SyncGroupLogTypeError   SyncGroupLogType = "Error"
	SyncGroupLogTypeSuccess SyncGroupLogType = "Success"
	SyncGroupLogTypeWarning SyncGroupLogType = "Warning"
)

func PossibleValuesForSyncGroupLogType() []string {
	return []string{
		string(SyncGroupLogTypeAll),
		string(SyncGroupLogTypeError),
		string(SyncGroupLogTypeSuccess),
		string(SyncGroupLogTypeWarning),
	}
}

func (s *SyncGroupLogType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncGroupLogType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncGroupLogType(input string) (*SyncGroupLogType, error) {
	vals := map[string]SyncGroupLogType{
		"all":     SyncGroupLogTypeAll,
		"error":   SyncGroupLogTypeError,
		"success": SyncGroupLogTypeSuccess,
		"warning": SyncGroupLogTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncGroupLogType(input)
	return &out, nil
}

type SyncGroupState string

const (
	SyncGroupStateError       SyncGroupState = "Error"
	SyncGroupStateGood        SyncGroupState = "Good"
	SyncGroupStateNotReady    SyncGroupState = "NotReady"
	SyncGroupStateProgressing SyncGroupState = "Progressing"
	SyncGroupStateWarning     SyncGroupState = "Warning"
)

func PossibleValuesForSyncGroupState() []string {
	return []string{
		string(SyncGroupStateError),
		string(SyncGroupStateGood),
		string(SyncGroupStateNotReady),
		string(SyncGroupStateProgressing),
		string(SyncGroupStateWarning),
	}
}

func (s *SyncGroupState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSyncGroupState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSyncGroupState(input string) (*SyncGroupState, error) {
	vals := map[string]SyncGroupState{
		"error":       SyncGroupStateError,
		"good":        SyncGroupStateGood,
		"notready":    SyncGroupStateNotReady,
		"progressing": SyncGroupStateProgressing,
		"warning":     SyncGroupStateWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncGroupState(input)
	return &out, nil
}

type Type string

const (
	TypeAll     Type = "All"
	TypeError   Type = "Error"
	TypeSuccess Type = "Success"
	TypeWarning Type = "Warning"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeAll),
		string(TypeError),
		string(TypeSuccess),
		string(TypeWarning),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"all":     TypeAll,
		"error":   TypeError,
		"success": TypeSuccess,
		"warning": TypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
