package integrationruntimes

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeAuthKeyName string

const (
	IntegrationRuntimeAuthKeyNameAuthKeyOne IntegrationRuntimeAuthKeyName = "authKey1"
	IntegrationRuntimeAuthKeyNameAuthKeyTwo IntegrationRuntimeAuthKeyName = "authKey2"
)

func PossibleValuesForIntegrationRuntimeAuthKeyName() []string {
	return []string{
		string(IntegrationRuntimeAuthKeyNameAuthKeyOne),
		string(IntegrationRuntimeAuthKeyNameAuthKeyTwo),
	}
}

func (s *IntegrationRuntimeAuthKeyName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationRuntimeAuthKeyName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationRuntimeAuthKeyName(input string) (*IntegrationRuntimeAuthKeyName, error) {
	vals := map[string]IntegrationRuntimeAuthKeyName{
		"authkey1": IntegrationRuntimeAuthKeyNameAuthKeyOne,
		"authkey2": IntegrationRuntimeAuthKeyNameAuthKeyTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeAuthKeyName(input)
	return &out, nil
}

type IntegrationRuntimeAutoUpdate string

const (
	IntegrationRuntimeAutoUpdateOff IntegrationRuntimeAutoUpdate = "Off"
	IntegrationRuntimeAutoUpdateOn  IntegrationRuntimeAutoUpdate = "On"
)

func PossibleValuesForIntegrationRuntimeAutoUpdate() []string {
	return []string{
		string(IntegrationRuntimeAutoUpdateOff),
		string(IntegrationRuntimeAutoUpdateOn),
	}
}

func (s *IntegrationRuntimeAutoUpdate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationRuntimeAutoUpdate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationRuntimeAutoUpdate(input string) (*IntegrationRuntimeAutoUpdate, error) {
	vals := map[string]IntegrationRuntimeAutoUpdate{
		"off": IntegrationRuntimeAutoUpdateOff,
		"on":  IntegrationRuntimeAutoUpdateOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeAutoUpdate(input)
	return &out, nil
}

type IntegrationRuntimeState string

const (
	IntegrationRuntimeStateAccessDenied     IntegrationRuntimeState = "AccessDenied"
	IntegrationRuntimeStateInitial          IntegrationRuntimeState = "Initial"
	IntegrationRuntimeStateLimited          IntegrationRuntimeState = "Limited"
	IntegrationRuntimeStateNeedRegistration IntegrationRuntimeState = "NeedRegistration"
	IntegrationRuntimeStateOffline          IntegrationRuntimeState = "Offline"
	IntegrationRuntimeStateOnline           IntegrationRuntimeState = "Online"
	IntegrationRuntimeStateStarted          IntegrationRuntimeState = "Started"
	IntegrationRuntimeStateStarting         IntegrationRuntimeState = "Starting"
	IntegrationRuntimeStateStopped          IntegrationRuntimeState = "Stopped"
	IntegrationRuntimeStateStopping         IntegrationRuntimeState = "Stopping"
)

func PossibleValuesForIntegrationRuntimeState() []string {
	return []string{
		string(IntegrationRuntimeStateAccessDenied),
		string(IntegrationRuntimeStateInitial),
		string(IntegrationRuntimeStateLimited),
		string(IntegrationRuntimeStateNeedRegistration),
		string(IntegrationRuntimeStateOffline),
		string(IntegrationRuntimeStateOnline),
		string(IntegrationRuntimeStateStarted),
		string(IntegrationRuntimeStateStarting),
		string(IntegrationRuntimeStateStopped),
		string(IntegrationRuntimeStateStopping),
	}
}

func (s *IntegrationRuntimeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationRuntimeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationRuntimeState(input string) (*IntegrationRuntimeState, error) {
	vals := map[string]IntegrationRuntimeState{
		"accessdenied":     IntegrationRuntimeStateAccessDenied,
		"initial":          IntegrationRuntimeStateInitial,
		"limited":          IntegrationRuntimeStateLimited,
		"needregistration": IntegrationRuntimeStateNeedRegistration,
		"offline":          IntegrationRuntimeStateOffline,
		"online":           IntegrationRuntimeStateOnline,
		"started":          IntegrationRuntimeStateStarted,
		"starting":         IntegrationRuntimeStateStarting,
		"stopped":          IntegrationRuntimeStateStopped,
		"stopping":         IntegrationRuntimeStateStopping,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeState(input)
	return &out, nil
}

type IntegrationRuntimeType string

const (
	IntegrationRuntimeTypeManaged    IntegrationRuntimeType = "Managed"
	IntegrationRuntimeTypeSelfHosted IntegrationRuntimeType = "SelfHosted"
)

func PossibleValuesForIntegrationRuntimeType() []string {
	return []string{
		string(IntegrationRuntimeTypeManaged),
		string(IntegrationRuntimeTypeSelfHosted),
	}
}

func (s *IntegrationRuntimeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIntegrationRuntimeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIntegrationRuntimeType(input string) (*IntegrationRuntimeType, error) {
	vals := map[string]IntegrationRuntimeType{
		"managed":    IntegrationRuntimeTypeManaged,
		"selfhosted": IntegrationRuntimeTypeSelfHosted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IntegrationRuntimeType(input)
	return &out, nil
}
