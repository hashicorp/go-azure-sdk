package serverautomatictuning

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

type AutomaticTuningServerMode string

const (
	AutomaticTuningServerModeAuto        AutomaticTuningServerMode = "Auto"
	AutomaticTuningServerModeCustom      AutomaticTuningServerMode = "Custom"
	AutomaticTuningServerModeUnspecified AutomaticTuningServerMode = "Unspecified"
)

func PossibleValuesForAutomaticTuningServerMode() []string {
	return []string{
		string(AutomaticTuningServerModeAuto),
		string(AutomaticTuningServerModeCustom),
		string(AutomaticTuningServerModeUnspecified),
	}
}

func (s *AutomaticTuningServerMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticTuningServerMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticTuningServerMode(input string) (*AutomaticTuningServerMode, error) {
	vals := map[string]AutomaticTuningServerMode{
		"auto":        AutomaticTuningServerModeAuto,
		"custom":      AutomaticTuningServerModeCustom,
		"unspecified": AutomaticTuningServerModeUnspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticTuningServerMode(input)
	return &out, nil
}

type AutomaticTuningServerReason string

const (
	AutomaticTuningServerReasonAutoConfigured AutomaticTuningServerReason = "AutoConfigured"
	AutomaticTuningServerReasonDefault        AutomaticTuningServerReason = "Default"
	AutomaticTuningServerReasonDisabled       AutomaticTuningServerReason = "Disabled"
)

func PossibleValuesForAutomaticTuningServerReason() []string {
	return []string{
		string(AutomaticTuningServerReasonAutoConfigured),
		string(AutomaticTuningServerReasonDefault),
		string(AutomaticTuningServerReasonDisabled),
	}
}

func (s *AutomaticTuningServerReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAutomaticTuningServerReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAutomaticTuningServerReason(input string) (*AutomaticTuningServerReason, error) {
	vals := map[string]AutomaticTuningServerReason{
		"autoconfigured": AutomaticTuningServerReasonAutoConfigured,
		"default":        AutomaticTuningServerReasonDefault,
		"disabled":       AutomaticTuningServerReasonDisabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AutomaticTuningServerReason(input)
	return &out, nil
}
