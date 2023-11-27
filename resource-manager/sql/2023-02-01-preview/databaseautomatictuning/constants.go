package databaseautomatictuning

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomaticTuningDisabledReason string

const (
	AutomaticTuningDisabledReasonAutoConfigured      AutomaticTuningDisabledReason = "AutoConfigured"
	AutomaticTuningDisabledReasonDefault             AutomaticTuningDisabledReason = "Default"
	AutomaticTuningDisabledReasonDisabled            AutomaticTuningDisabledReason = "Disabled"
	AutomaticTuningDisabledReasonInheritedFromServer AutomaticTuningDisabledReason = "InheritedFromServer"
	AutomaticTuningDisabledReasonNotSupported        AutomaticTuningDisabledReason = "NotSupported"
	AutomaticTuningDisabledReasonQueryStoreOff       AutomaticTuningDisabledReason = "QueryStoreOff"
	AutomaticTuningDisabledReasonQueryStoreReadOnly  AutomaticTuningDisabledReason = "QueryStoreReadOnly"
)

func PossibleValuesForAutomaticTuningDisabledReason() []string {
	return []string{
		string(AutomaticTuningDisabledReasonAutoConfigured),
		string(AutomaticTuningDisabledReasonDefault),
		string(AutomaticTuningDisabledReasonDisabled),
		string(AutomaticTuningDisabledReasonInheritedFromServer),
		string(AutomaticTuningDisabledReasonNotSupported),
		string(AutomaticTuningDisabledReasonQueryStoreOff),
		string(AutomaticTuningDisabledReasonQueryStoreReadOnly),
	}
}

func (s *AutomaticTuningDisabledReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticTuningDisabledReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticTuningDisabledReason(input string) (*AutomaticTuningDisabledReason, error) {
	vals := map[string]AutomaticTuningDisabledReason{
		"autoconfigured":      AutomaticTuningDisabledReasonAutoConfigured,
		"default":             AutomaticTuningDisabledReasonDefault,
		"disabled":            AutomaticTuningDisabledReasonDisabled,
		"inheritedfromserver": AutomaticTuningDisabledReasonInheritedFromServer,
		"notsupported":        AutomaticTuningDisabledReasonNotSupported,
		"querystoreoff":       AutomaticTuningDisabledReasonQueryStoreOff,
		"querystorereadonly":  AutomaticTuningDisabledReasonQueryStoreReadOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticTuningDisabledReason(input)
	return &out, nil
}

type AutomaticTuningMode string

const (
	AutomaticTuningModeAuto        AutomaticTuningMode = "Auto"
	AutomaticTuningModeCustom      AutomaticTuningMode = "Custom"
	AutomaticTuningModeInherit     AutomaticTuningMode = "Inherit"
	AutomaticTuningModeUnspecified AutomaticTuningMode = "Unspecified"
)

func PossibleValuesForAutomaticTuningMode() []string {
	return []string{
		string(AutomaticTuningModeAuto),
		string(AutomaticTuningModeCustom),
		string(AutomaticTuningModeInherit),
		string(AutomaticTuningModeUnspecified),
	}
}

func (s *AutomaticTuningMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticTuningMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticTuningMode(input string) (*AutomaticTuningMode, error) {
	vals := map[string]AutomaticTuningMode{
		"auto":        AutomaticTuningModeAuto,
		"custom":      AutomaticTuningModeCustom,
		"inherit":     AutomaticTuningModeInherit,
		"unspecified": AutomaticTuningModeUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticTuningMode(input)
	return &out, nil
}

type AutomaticTuningOptionModeActual string

const (
	AutomaticTuningOptionModeActualOff AutomaticTuningOptionModeActual = "Off"
	AutomaticTuningOptionModeActualOn  AutomaticTuningOptionModeActual = "On"
)

func PossibleValuesForAutomaticTuningOptionModeActual() []string {
	return []string{
		string(AutomaticTuningOptionModeActualOff),
		string(AutomaticTuningOptionModeActualOn),
	}
}

func (s *AutomaticTuningOptionModeActual) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticTuningOptionModeActual(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticTuningOptionModeActual(input string) (*AutomaticTuningOptionModeActual, error) {
	vals := map[string]AutomaticTuningOptionModeActual{
		"off": AutomaticTuningOptionModeActualOff,
		"on":  AutomaticTuningOptionModeActualOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticTuningOptionModeActual(input)
	return &out, nil
}

type AutomaticTuningOptionModeDesired string

const (
	AutomaticTuningOptionModeDesiredDefault AutomaticTuningOptionModeDesired = "Default"
	AutomaticTuningOptionModeDesiredOff     AutomaticTuningOptionModeDesired = "Off"
	AutomaticTuningOptionModeDesiredOn      AutomaticTuningOptionModeDesired = "On"
)

func PossibleValuesForAutomaticTuningOptionModeDesired() []string {
	return []string{
		string(AutomaticTuningOptionModeDesiredDefault),
		string(AutomaticTuningOptionModeDesiredOff),
		string(AutomaticTuningOptionModeDesiredOn),
	}
}

func (s *AutomaticTuningOptionModeDesired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticTuningOptionModeDesired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticTuningOptionModeDesired(input string) (*AutomaticTuningOptionModeDesired, error) {
	vals := map[string]AutomaticTuningOptionModeDesired{
		"default": AutomaticTuningOptionModeDesiredDefault,
		"off":     AutomaticTuningOptionModeDesiredOff,
		"on":      AutomaticTuningOptionModeDesiredOn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticTuningOptionModeDesired(input)
	return &out, nil
}
